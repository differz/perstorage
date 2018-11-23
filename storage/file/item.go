package file

import (
	"log"
	"os"

	"../../core"
	"github.com/satori/go.uuid"
)

// StoreItem save file to storage
func (s Storage) StoreItem(item core.Item) (int, error) {
	ns, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8") // TODO: namespace
	if err != nil {
		log.Printf("can't generate namespace %e", err)
	}
	key := "salt:" + item.Filename
	u5 := uuid.NewV5(ns, key)
	if err != nil {
		log.Printf("can't generate UUID %e", err)
	}

	dir := s.dir + u5.String() + "/"
	path := dir + item.Filename

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Printf("can't create directory %s %e", dir, err)
	}
	err = os.Rename(item.SourceName, path)
	if err != nil {
		log.Printf("can't rename file %s to %s %e", item.SourceName, path, err)
	}

	fi, err := os.Stat(path)
	size := int64(-1)
	if err == nil {
		size = fi.Size()
	}

	mutex.Lock()
	defer mutex.Unlock()

	sql := "INSERT INTO items(name, filename, path, size, category, available) VALUES(?, ?, ?, ?, ?, ?)"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal("prepared statement for table items ", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("", item.Filename, path, size, item.Category, true)
	if err != nil {
		log.Printf("can't insert item into db %e", err)
		return 0, err
	}

	if item.IsNew() {
		id, err := res.LastInsertId()
		if err != nil {
			log.Printf("error get last inserted id for item %e", err)
		}
		item.ID = int(id)
	}
	return item.ID, err
}

// FindItemByID get file from storage
func (s Storage) FindItemByID(id int) (core.Item, bool) {
	return core.Item{}, false
}

// DeleteItem remove file from storage
func (s Storage) DeleteItem(item core.Item) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if s.deleteItemFile(item) {
		return s.deleteItemLink(item)
	}
	return false
}

func (s Storage) deleteItemFile(item core.Item) bool {
	sql := "SELECT path FROM items WHERE id = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(item.ID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if rows.Next() {
		var path string
		err = rows.Scan(&path)
		if err != nil {
			log.Fatal(err)
		}
		err = os.Remove(path)
		if err != nil {
			log.Print(err)
			return false
		}
	}
	return true
}

func (s Storage) deleteItemLink(item core.Item) bool {
	sql := "DELETE FROM items WHERE id = ?"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal("prepared statement for table items ", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(item.ID)
	if err != nil {
		log.Fatal("can't delete item from db ", err)
	}
	return true
}

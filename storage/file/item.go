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

	sql := "INSERT INTO items(name, filename, path, size, available) VALUES(?, ?, ?, ?, ?)"
	stmt, err := s.connection.Prepare(sql)
	if err != nil {
		log.Fatal("prepared statement for table items ", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec("", item.Filename, path, size, true)
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

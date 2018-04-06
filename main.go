package main

import (
	"fmt"
	"net/http"

	"./storage"
	_ "./storage/filestorage"
	_ "./storage/mongostorage"
	"./web"
	_ "./web/upload"
	// https://help.compose.com/docs/mongodb-go-compose
)

var result []struct {
	ID    int //"_id"
	Value int
}

func main() {
	fmt.Println("Start...")
	storage.Print()
	web.Print()

	//	db, _ := storage.Get("file", 1)
	//	fmt.Println(db)

	//	con, _ := web.Get("api", 1)
	//	fmt.Println(con)

	http.ListenAndServe(":8081", nil)
}

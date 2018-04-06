package main

import (
	"fmt"
	"net/http"

	"./storage"
	_ "./storage/filestorage"
	_ "./storage/mongostorage"
	"./web"
	_ "./web/root"
	_ "./web/upload"
)

func main() {
	fmt.Println("Start...")
	storage.Print()
	web.Print()

	db, _ := storage.Get("file", 1)
	fmt.Println(db)

	con, _ := web.Get("root", 1)
	fmt.Println(con)

	http.ListenAndServe(":8081", nil)
}

package main

import (
	"fmt"
	"net/http"

	"./configuration"
	"./storage"
	_ "./storage/filestorage"
	_ "./storage/mongostorage"
	"./web"
	_ "./web/download"
	_ "./web/root"
	_ "./web/upload"
)

func main() {
	fmt.Println("Start...")
	storage.Print()
	web.Print()

	cfg := configuration.Get()
	defer cfg.Close()

	con, _ := web.Get("root", 1)
	fmt.Println(con)

	http.ListenAndServe(":8081", nil)
}

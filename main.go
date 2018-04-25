package main

import (
	"fmt"
	"net/http"

	"./configuration"
	"./messenger"
	"./messenger/service"
	"./storage"
	"./web"

	// used modules
	_ "./messenger/telegram"
	_ "./storage/filestorage"
	_ "./storage/mongostorage"
	_ "./web/download"
	_ "./web/root"
	_ "./web/upload"
)

func main() {
	defer configuration.Close()

	fmt.Println("Start...")
	messenger.Print()
	storage.Print()
	web.Print()

	con, _ := web.Get("root", 1)
	fmt.Println(con)

	srv := messengers.NewService()
	srv.ListenChat()

	http.ListenAndServe(":8081", nil)
}

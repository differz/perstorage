package main

import (
	"net/http"

	"./common"
	"./configuration/context"
	"./messenger"
	"./messenger/service"
	"./storage"
	"./web"

	// used modules
	_ "./messenger/telegram"
	_ "./storage/file"
	_ "./storage/mongo"
	_ "./web/download"
	_ "./web/root"
	_ "./web/upload"
)

func main() {
	defer context.Close()

	common.ContextUpMessage("application", "start...")
	messenger.Print()
	storage.Print()
	web.Print()

	srv := messengers.NewService()
	srv.ListenChat()

	http.ListenAndServe(":8081", nil)
}

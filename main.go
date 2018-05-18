package main

import (
	"log"
	"net/http"

	"./common"
	"./configuration"
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

	common.ContextUpMessage("application", "running...")
	err := http.ListenAndServe(configuration.ServerPort(), nil)
	if err != nil {
		log.Printf("can't start http server %e", err)
	}
}

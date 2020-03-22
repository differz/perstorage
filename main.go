package main

import (
	"log"
	"net/http"

	"perstorage/common"
	"perstorage/configuration"
	"perstorage/configuration/context"
	"perstorage/messenger"
	"perstorage/messenger/messengers"

	"perstorage/storage"
	"perstorage/web"

	// used modules
	_ "perstorage/messenger/telegram"
	_ "perstorage/storage/file"
	_ "perstorage/storage/mongo"
	_ "perstorage/web/api"
	_ "perstorage/web/download"
	_ "perstorage/web/root"
	_ "perstorage/web/upload"
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

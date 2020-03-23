package main

import (
	"log"
	"net/http"

	"github.com/differz/perstorage/common"
	"github.com/differz/perstorage/configuration"
	"github.com/differz/perstorage/configuration/context"
	"github.com/differz/perstorage/messenger"
	"github.com/differz/perstorage/messenger/messengers"

	"github.com/differz/perstorage/storage"
	"github.com/differz/perstorage/web"

	// used modules
	_ "github.com/differz/perstorage/messenger/telegram"
	_ "github.com/differz/perstorage/storage/file"
	_ "github.com/differz/perstorage/storage/mongo"
	_ "github.com/differz/perstorage/web/api"
	_ "github.com/differz/perstorage/web/download"
	_ "github.com/differz/perstorage/web/root"
	_ "github.com/differz/perstorage/web/upload"
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
	err := http.ListenAndServeTLS(configuration.ServerPort(), "cert.pem", "privkey.pem", nil)
	if err != nil {
		log.Printf("can't start https server %e", err)
	}
}

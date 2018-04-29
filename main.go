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
	_ "./storage/filestorage"
	_ "./storage/mongostorage"
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

	//con, _ := web.Get("root")
	//fmt.Println(con)

	srv := messengers.NewService()
	srv.ListenChat()

	http.ListenAndServe(":8081", nil)
}

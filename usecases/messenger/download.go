package usecases

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/differz/perstorage/common"
	"github.com/differz/perstorage/configuration"
	"github.com/differz/perstorage/contracts/usecases"
)

// Download file from chat
func (r CustomerMessengerResponse) downloadFile(url, filename string, size int, msgr string, chatID int, description string) {
	customerID, ok := r.repo.IsRegisteredChatID(chatID, msgr)
	if !ok {
		log.Printf("can't get customer by chatID %d", chatID)
		return
	}

	inMD5 := common.ConvertToMD5(url)
	dir := configuration.TempDirectory() + inMD5 + "/"

	req := contracts.PlaceOrderRequest{}
	req.Filename = filename
	req.Dir = dir
	req.CustomerID = customerID
	req.Description = description
	req.Private = true

	err := os.MkdirAll(req.Dir, os.ModePerm)
	if err != nil {
		log.Printf("can't create directory %s %e", req.Dir, err)
		return
	}

	temp, err := os.Create(req.GetFullFileName())
	if err != nil {
		log.Printf("can't create file %e", err)
		return
	}

	page, err := http.Get(url)
	if err != nil {
		log.Printf("can't get URL %e", err)
		return
	}
	defer page.Body.Close()

	_, err = io.Copy(temp, page.Body)
	if err != nil {
		log.Printf("can't copy from URL %e", err)
		return
	}

	req.MD5 = common.ComputeFileMD5(temp)
	temp.Close()

	resp := PlaceOrderResponse{phone: req.Phone, orderMessage: r.orderMessage, description: description}
	r.placeOrder.PlaceOrder(req, resp)
}

// PlaceOrderResponse response data
type PlaceOrderResponse struct {
	downloadLink string
	description  string
	phone        string
	orderMessage contracts.OrderMessageInput
}

// OnResponse send order message through messenger via registered phone number
func (r PlaceOrderResponse) OnResponse(phone, orderLink, description string) {
	r.downloadLink = orderLink
	r.description = description ///+?
	r.phone = phone

	r.orderMessage.OrderMessage(phone, orderLink, description)
}

package usecases

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"os"

	"../contracts/messengers"
	"../contracts/usecases"
	"../core"
	"../messenger"
	"../storage"
)

// CustomerMessengerUseCase object
type CustomerMessengerUseCase struct {
	repo         storage.Storager
	msgr         messenger.Messenger
	placeOrder   contracts.PlaceOrderInput
	orderMessage contracts.OrderMessageInput
	//
	subject     string
	description string
}

// NewCustomerMessengerUseCase constructor
func NewCustomerMessengerUseCase(repo storage.Storager, msgr messenger.Messenger) CustomerMessengerUseCase {
	return CustomerMessengerUseCase{
		repo:         repo,
		msgr:         msgr,
		placeOrder:   NewPlaceOrderUseCase(repo),
		orderMessage: NewOrderMessageUseCase(repo, msgr),
		description:  "new customer messenger",
	}
}

// ListenChat listen messengers chat
func (u CustomerMessengerUseCase) ListenChat() {
	request := messengers.ListenChatRequest{
		Repo: u.repo,
	}
	output := CustomerMessengerResponse{
		repo:         u.repo,
		msgr:         u.msgr,
		placeOrder:   u.placeOrder,
		orderMessage: u.orderMessage,
	}
	go u.msgr.ListenChat(request, output)
}

// CustomerMessengerResponse object to response
type CustomerMessengerResponse struct {
	repo         storage.Storager
	msgr         messenger.Messenger
	placeOrder   contracts.PlaceOrderInput
	orderMessage contracts.OrderMessageInput
}

// OnResponse register new chatID to customer messenger
func (r CustomerMessengerResponse) OnResponse(request messengers.ListenChatRequest) {
	if validatePhone(request.Phone) {
		go r.registerMessenger(request.Phone, request.Messenger, request.ChatID)
	}
	if request.FileURL != "" {
		go r.downloadFile(request.FileURL, request.FileName, request.FileSize, request.Messenger, request.ChatID)
	}
}

func (r CustomerMessengerResponse) registerMessenger(phone, messenger string, chatID int) {
	customerID, err := core.GetCustomerIDByPhone(phone)
	if err != nil {
		log.Printf("can't get customer id by phone %s %e", phone, err)
		return
	}

	customer, ok := r.repo.FindCustomerByID(customerID)
	if !ok {
		customer.ID = customerID
		customer.Phone = phone
		r.repo.StoreCustomer(customer)
	}

	_, ok = r.repo.FindCustomerChatID(customer, messenger)
	if !ok {
		r.repo.StoreCustomerMessenger(customer, messenger, chatID)
	}
}

func validatePhone(phone string) bool {
	return phone != ""
}

// TODO REF

func (r CustomerMessengerResponse) downloadFile(url, name string, size int, msgr string, chatID int) {
	// TODO

	customerID, ok := r.repo.IsRegisteredChatID(chatID, msgr)
	if ok {
		r.DownloadFile(url, name, customerID)
	}
}

// DownloadFile from chat
func (r CustomerMessengerResponse) DownloadFile(url, filename string, customerID int) {

	// TODO ref
	hasher := md5.New()
	hasher.Write([]byte(url))
	hash := hasher.Sum(nil)
	inMD5 := hex.EncodeToString(hash)

	req := contracts.PlaceOrderRequest{}
	req.Filename = filename
	req.Dir = "./local/incoming/" + inMD5 + "/"
	req.CustomerID = customerID
	req.Private = true

	err := os.MkdirAll(req.Dir, os.ModePerm)
	if err != nil {
		log.Printf("can't create directory %s %e", req.Dir, err)
		return
	}

	temp, err := os.Create(req.GetSourceName())
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

	req.MD5 = computeMD5(temp)
	temp.Close()

	resp := PlaceOrderResponse{phone: req.Phone, orderMessage: r.orderMessage}
	r.placeOrder.PlaceOrder(req, resp)

	//return resp.downloadLink, nil

}
func computeMD5(file *os.File) []byte {
	var result []byte
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result
	}
	return hash.Sum(result)
}

// PlaceOrderResponse response data
type PlaceOrderResponse struct {
	downloadLink string
	phone        string
	orderMessage contracts.OrderMessageInput
}

// OnResponse send order message through messenger via registered phone number
func (r PlaceOrderResponse) OnResponse(phone, orderLink string) {
	r.downloadLink = orderLink
	r.phone = phone

	r.orderMessage.OrderMessage(phone, orderLink)
}

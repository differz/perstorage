package messengers

import (
	"perstorage/storage"
)

// ListenChatRequest parameters
type ListenChatRequest struct {
	Repo        storage.Storager
	Phone       string
	Messenger   string
	ChatID      int
	FileURL     string
	FileName    string
	FileSize    int
	Description string
}

// ListenChatInput listen chat input contract
type ListenChatInput interface {
	ListenChat(request ListenChatRequest, output ListenChatOutput)
}

// ListenChatOutput listen chat output contract
type ListenChatOutput interface {
	OnResponse(request ListenChatRequest)
}

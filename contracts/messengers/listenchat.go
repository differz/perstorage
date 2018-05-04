package messengers

import (
	"../../storage"
)

// ListenChatRequest parameters
type ListenChatRequest struct {
	Repo      storage.Storager
	Phone     string
	Messenger string
	ChatID    int
	FileID    string
	FileName  string
	FileSize  int
}

// ListenChatInput listen chat input contract
type ListenChatInput interface {
	ListenChat(request ListenChatRequest, output ListenChatOutput)
}

// ListenChatOutput listen chat output contract
type ListenChatOutput interface {
	OnResponse(request ListenChatRequest)
}

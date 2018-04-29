package messengers

// ListenChatInput listen chat input contract
type ListenChatInput interface {
	ListenChat(output ListenChatOutput)
}

// ListenChatOutput listen chat output contract
type ListenChatOutput interface {
	OnResponse(phone string, messengerName string, chatID int)
}

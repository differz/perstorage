package messengers

// ListenChat ...
type ListenChatInput interface {
	ListenChat(output ListenChatOutput)
}

type ListenChatOutput interface {
	OnResponse(phone string, messengerName string, chatID int)
}

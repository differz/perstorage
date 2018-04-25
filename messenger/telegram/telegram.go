package telegram

import (
	"fmt"
	"log"

	"../../contracts/messengers"
	"../../messenger"
	"gopkg.in/telegram-bot-api.v4"
)

// Messenge ...
type Messenge struct {
	name string
	bot  *tgbotapi.BotAPI
}

// New create instance. Init method has pointer receiver
func New() *Messenge {
	return &Messenge{
		name: "telegram",
	}
}

// Init ...
func (m *Messenge) Init(args ...string) {
	fmt.Println("Init telegram")
	token := args[0]
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	m.bot = bot
}

// ListenChat ...
func (m Messenge) ListenChat(output messengers.ListenChatOutput) {
	bot := m.bot
	tgu := tgbotapi.NewUpdate(0)
	tgu.Timeout = 60

	updates, err := bot.GetUpdatesChan(tgu)
	if err != nil {
		log.Panic(err)
	}
	// TODO: if chatId is registered
	for update := range updates {
		chatID := update.Message.Chat.ID
		if update.Message.Contact != nil {
			phone := update.Message.Contact.PhoneNumber
			output.OnResponse(phone, m.name, int(chatID))
			continue
		}

		msg := tgbotapi.NewMessage(chatID, "send you phone to register")
		var keyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButtonContact("\xF0\x9F\x93\x9E Send phone"),
			),
		)
		msg.ReplyMarkup = keyboard
		bot.Send(msg)
	}
}

// ShowOrder ...
func (m Messenge) ShowOrder(chatID int, message string) error {
	// TODO: add server
	downloadLink := "http://localhost:8081/download/" + message
	msg := tgbotapi.NewMessage(int64(chatID), downloadLink)
	m.bot.Send(msg)
	return nil
}

func init() {
	messenger.Register("telegram", New())
}

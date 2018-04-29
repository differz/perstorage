package telegram

import (
	"fmt"
	"log"

	"../../common"
	"../../contracts/messengers"
	"../../messenger"
	"gopkg.in/telegram-bot-api.v4"
)

// Messenge telegram object
type Messenge struct {
	name string
	bot  *tgbotapi.BotAPI
}

const component = "telegram"

// New create instance. Init method has pointer receiver
func New() *Messenge {
	return &Messenge{
		name: "telegram",
	}
}

// Init connect to API by token
func (m *Messenge) Init(args ...string) error {
	common.ContextUpMessage(component, "init telegram token")
	token := args[0]
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("can't connect with token ", token, " ", err)
	} else {
		m.bot = bot
	}
	return err
}

// Name telegram
func (m Messenge) Name() string {
	return m.name
}

// Available true if telegram api is on
func (m Messenge) Available() bool {
	return m.bot != nil
}

// ListenChat send all new messages to output interface
func (m Messenge) ListenChat(output messengers.ListenChatOutput) {
	if !m.Available() {
		log.Println("bot not available")
		return
	}
	bot := m.bot
	tgu := tgbotapi.NewUpdate(0)
	tgu.Timeout = 60

	updates, err := bot.GetUpdatesChan(tgu)
	if err != nil {
		log.Panic("can't get updates chanel ", err)
	}

	for update := range updates {
		chatID := update.Message.Chat.ID // TODO: if chatId is registered
		// is contact?
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

// ShowOrder place order details in chat message
func (m Messenge) ShowOrder(chatID int, message string) error {
	if !m.Available() {
		return fmt.Errorf("bot not available")
	}
	// TODO: add server
	downloadLink := "http://localhost:8081/download/" + message
	msg := tgbotapi.NewMessage(int64(chatID), downloadLink)
	_, err := m.bot.Send(msg)
	return err
}

func (m Messenge) String() string {
	return m.name
}

func init() {
	messenger.Register(component, New())
}

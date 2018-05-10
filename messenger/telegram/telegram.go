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
	name   string
	server string
	bot    *tgbotapi.BotAPI
}

const (
	component = "telegram"
	register  = "send you phone to register"
	upload    = "<b>to upload files</b>" + "\n" + "<i>use Send as File!</i>" + "\n" + "<em>for > 1,5G follow link:</em>"
)

// New create instance. Init method has pointer receiver
func New() *Messenge {
	return &Messenge{
		name: "telegram",
	}
}

// Init connect to API by token
func (m *Messenge) Init(args ...string) error {
	common.ContextUpMessage(component, "init telegram token")
	m.server = args[1]
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
func (m Messenge) ListenChat(request messengers.ListenChatRequest, output messengers.ListenChatOutput) {
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
		chatID := update.Message.Chat.ID
		// is contact?
		con := update.Message.Contact
		if con != nil {
			request.Phone = con.PhoneNumber
			request.Messenger = m.name
			request.ChatID = int(chatID)
			output.OnResponse(request)
		}
		// is document?
		doc := update.Message.Document
		if doc != nil {
			url, err := m.bot.GetFileDirectURL(doc.FileID)
			if err != nil {
				log.Printf("can't get url for file from chat %e", err)

			} else {
				request.Messenger = m.name
				request.ChatID = int(chatID)
				request.FileURL = url
				request.FileName = doc.FileName
				request.FileSize = doc.FileSize
				output.OnResponse(request)
			}
		}

		_, registered := request.Repo.IsRegisteredChatID(int(chatID), m.name)
		if registered {
			msg := tgbotapi.NewMessage(chatID, upload)
			msg.ParseMode = "HTML"
			bot.Send(msg)
			msg = tgbotapi.NewMessage(chatID, m.server)
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(chatID, register)
			var keyboard = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButtonContact("\xF0\x9F\x93\x9E Send phone"),
				),
			)
			msg.ReplyMarkup = keyboard
			bot.Send(msg)
		}
	}
}

// ShowOrder place order details in chat message
func (m Messenge) ShowOrder(chatID int, message string) error {
	if !m.Available() {
		return fmt.Errorf("bot not available")
	}
	downloadLink := common.DownloadLink(message)
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

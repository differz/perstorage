package telegram

import (
	"log"

	"../../common"
	"../../messenger"
	"gopkg.in/telegram-bot-api.v4"
)

// Messenge telegram object
type Messenge struct {
	name   string
	server string
	bot    *tgbotapi.BotAPI
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

func (m Messenge) String() string {
	return m.name
}

func init() {
	messenger.Register(component, New())
}

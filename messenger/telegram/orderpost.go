package telegram

import (
	"fmt"

	"github.com/differz/perstorage/common"

	"gopkg.in/telegram-bot-api.v4"
)

// ShowOrder place order details in chat message
func (m Messenge) ShowOrder(chatID int, message, description string) error {
	if !m.Available() {
		return fmt.Errorf("bot not available")
	}
	if description != "" {
		msg := tgbotapi.NewMessage(int64(chatID), description)
		_, err := m.bot.Send(msg)
		if err != nil {
			return err
		}
	}
	downloadLink := common.DownloadLink(message)
	msg := tgbotapi.NewMessage(int64(chatID), downloadLink)
	_, err := m.bot.Send(msg)
	return err
}

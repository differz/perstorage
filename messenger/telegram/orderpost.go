package telegram

import (
	"fmt"

	"../../common"
	"gopkg.in/telegram-bot-api.v4"
)

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

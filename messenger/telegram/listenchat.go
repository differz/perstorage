package telegram

import (
	"log"

	"../../contracts/messengers"
	"gopkg.in/telegram-bot-api.v4"
)

const (
	register = "send you phone to register"
	upload   = "<b>to upload files</b>" + "\n" + "<i>use Send as File!</i>" + "\n" + "<em>for > 1,5G follow link:</em>"
)

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

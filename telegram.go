package main

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

type user struct {
	last_name  string
	first_name string
	chat_id    int64
}

var (
	bot *tgbotapi.BotAPI
)

func bot_go() {
	var err error
	bot, err = tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		codename := update.Message.From.FirstName + "_" + update.Message.From.LastName
		newuser := user{last_name: update.Message.From.LastName, first_name: update.Message.From.FirstName, chat_id: update.Message.Chat.ID}
		users[codename] = newuser
		err := saveuser(newuser)
		if err != nil {
			log.Panic(err)
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
func sendmsg(m msg) error {
	user1, ok := users[m.rec]
	if ok == false {
		return fmt.Errorf("no user : %s", m.rec)
	}
	msg := tgbotapi.NewMessage(user1.chat_id, m.text)
	bot.Send(msg)
	return nil
}

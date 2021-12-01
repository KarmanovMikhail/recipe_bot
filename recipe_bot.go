package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	//tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type resStruct struct {
	Err                   string
	Result                string
	CacheUse              int
	Source                string
	From                  string
	SourceTransliteration string
	TargetTransliteration string
}

func main() {

	var tokenTGBot string
	const tokenStringTg = "tgkey"

	flag.StringVar(&tokenTGBot, tokenStringTg, "", "The token of your telegramm bot")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "About: Telegram bot with recipes for breackfast\n")
		fmt.Fprintf(os.Stderr, "Author: Karmanov Mikhail (karmanov_mihail@mail.ru)\n")
		fmt.Fprintf(os.Stderr, "Version: 0.0.1\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if tokenTGBot == "" {
		log.Panic("Please insert token of your telegramm bot!")
	}

	bot, err := tgbotapi.NewBotAPI(tokenTGBot)
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
		textInMessage := update.Message.Text
		myText := "234567890" + textInMessage
		//myText := sendToTranslater(tokenLgvnx, textInMessage)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, myText)

		bot.Send(msg)
	}
}

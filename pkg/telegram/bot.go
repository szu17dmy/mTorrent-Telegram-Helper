package telegram

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/telebot.v3"
)

const (
	EnvKeyTelegramBotToken  = "TELEGRAM_BOT_TOKEN"
	EnvKeyTelegramChannelID = "TELEGRAM_CHANNEL_ID"
)

var (
	bot     *telebot.Bot
	token   string
	channel telebot.ChatID
)

func SendString(message string) (*telebot.Message, error) {
	return bot.Send(channel, message)
}

func SendHtml(message string) (*telebot.Message, error) {
	return bot.Send(channel, message, &telebot.SendOptions{
		ParseMode: telebot.ModeHTML,
	})
}

func init() {
	initChatID()
	initTelebot()
}

func initTelebot() {
	tk := os.Getenv(EnvKeyTelegramBotToken)
	if tk == "" {
		log.Fatalln("failed to load telegram bot token.")
	}
	token = tk
	tb, err := telebot.NewBot(telebot.Settings{
		Token: token,
	})
	if err != nil {
		log.Fatalln("failed to create telebot instance.", err)
	}
	bot = tb
	go bot.Start()
}

func initChatID() {
	cid := os.Getenv(EnvKeyTelegramChannelID)
	if cid == "" {
		log.Fatalln("failed to load telegram channel id.")
	}
	id, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		log.Fatalln("failed to parse telegram channel id.", err)
	}
	channel = telebot.ChatID(id)
}

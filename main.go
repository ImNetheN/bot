package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *tgbotapi.BotAPI

func getToken() string {
	file, err := os.Open("token")
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		panic(err)
	}

	token := string(buf[:n-1])

	return token
}

func main() {
	var err error
	bot, err = tgbotapi.NewBotAPI(getToken())
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	updates := bot.GetUpdatesChan(u)

	go receiveUpdates(ctx, updates)

	fmt.Println("listening for updates, press enter to stop")

	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()
}

func receiveUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for {
		select {
		case <-ctx.Done():
			return
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update tgbotapi.Update) {
	switch {
	case update.Message != nil:
		handleMessage(update.Message)
	case update.CallbackQuery != nil:
		panic("not implemented yet")
	}
}

func handleMessage(message *tgbotapi.Message) {
	response := tgbotapi.NewMessage(message.Chat.ID, "пошел нафик я пока ниче не умею")
	bot.Send(response)
}

package main

import (
	"log"
	"os"

	bot "github.com/chamaloown/difus/Bot"
	godotenv "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bot.BotToken = os.Getenv("BOT_TOKEN")
	bot.Run()
}

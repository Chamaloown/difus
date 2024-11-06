package main

import (
	"fmt"
	"os"

	bot "github.com/chamaloown/difus/Bot"
	godotenv "github.com/joho/godotenv"
)
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error!")
	}

 	bot.BotToken = os.Getenv("BOT_TOKEN")
 	bot.Run()
}
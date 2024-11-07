package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	almanax "github.com/chamaloown/difus/Almanax"
)

var BotToken string

func Run() {
	discord, err := discordgo.New("Bot " + BotToken)

	if err != nil {
		log.Fatal(err)
	}

	discord.AddHandler(newMessage)

	discord.Open()
	
	go almanax.Run(discord)

	defer discord.Close()

	fmt.Println("Bot running....")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
  		return
	}

	if message.ChannelID == os.Getenv("CHANNEL_ID") {

		switch {
		case strings.Contains(message.Content, "!author"):
			discord.ChannelMessageSend(message.ChannelID, "Malo Landemaine")
		case strings.Contains(message.Content, "!help"):
			discord.ChannelMessageSend(message.ChannelID, "Hello World😃")
		case strings.Contains(message.Content, "!bye"):
			discord.ChannelMessageSend(message.ChannelID, "Good Bye👋")
		case strings.Contains(message.Content, "!alma"):
			discord.ChannelMessageSend(message.ChannelID, "Almanax")
		default:
		}
	}
}
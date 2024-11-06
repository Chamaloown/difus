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

func checkNilErr(e error) {
 if e != nil {
  log.Fatal("Error message")
 }
}

func Run() {

 // create a session
 discord, err := discordgo.New("Bot " + BotToken)
 checkNilErr(err)

 go almanax.Run()

 // add a event handler
 discord.AddHandler(newMessage)

 // Cronjob


 // open session
 discord.Open()
 defer discord.Close() // close session, after function termination

 // keep bot running untill there is NO os interruption (ctrl + C)
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
			discord.ChannelMessageSend(message.Content, "Malo Landemaine")
		case strings.Contains(message.Content, "!help"):
			discord.ChannelMessageSend("314480237817626624", "Hello World😃")
		case strings.Contains(message.Content, "!bye"):
			discord.ChannelMessageSend("314480237817626624", "Good Bye👋")
		default:
		}
	}
}
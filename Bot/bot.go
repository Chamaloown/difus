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

	switch {
		case strings.Contains(message.Content, "!author"):
			discord.ChannelMessageSend(message.ChannelID, "Malo Landemaine")
		case strings.Contains(message.Content, "!help"):
			discord.ChannelMessageSend(message.ChannelID, help())
		case strings.Contains(message.Content, "!alma"): 
			var msg = almanax.GetAlmanax(message.Content)
			discord.ChannelMessageSendComplex(message.ChannelID, &msg)
		default:
	}
}

func help() string {
	return `Voici les commandes disponibles :

	📜 **!author** - Affiche le nom de l'auteur.
	❓ **!help** - Affiche ce message d'aide.
	📅 **!alma [today | week | JJ/MM/AAAA]** - Récupère l'Almanax pour un jour spécifique :
	      •  **today** : Affiche l'Almanax d'aujourd'hui.
	      •  **week** : Affiche l'Almanax pour toute la semaine.
	      •  **JJ/MM/AAAA** : Affiche l'Almanax pour une date spécifique (ex. 08/11/2024).
	
	Veuillez utiliser le bon format de date ou les mots-clés spécifiés pour chaque option.`

	}
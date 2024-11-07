package almanax

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	db "github.com/chamaloown/difus/Almanax/Db"
	gocron "github.com/go-co-op/gocron/v2"
)

func formatAlmanax(almanax db.Almanax) string {
	return fmt.Sprintf(
		"@here Salut les Dofusiens !\n\n📅 Almanax du **%s**\n\n🔮 **Méryde** : %s\n📈 **Type de Bonus** : %s\n🎁 **Bonus** : %s\n🎒 **Offrande** : %s x%d\n💰 **Prix estimé** : %d kamas\n",
		almanax.Date.Format("2006/01/02"), almanax.Merydes, almanax.Type, almanax.Bonus, almanax.Offerings, almanax.QuantityOffered, almanax.Kamas)
}

func Run(discord *discordgo.Session) {
	fmt.Println("Loading almanax...")
	db.Setup()
	fmt.Println("Almanax Loaded!")

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	j, err := s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(0, 36, 0),
				gocron.NewAtTime(22, 0, 0),
			),
		),
		gocron.NewTask(
			func() {
				pg := db.GetDBInstance()
				alamanax, err := db.GetAlmanax(pg, time.Now().AddDate(1, 0, 0))
				if err != nil {
					log.Fatal(err)
				}
				var message = discordgo.MessageSend{
					Content: formatAlmanax(alamanax),
				}
				discord.ChannelMessageSendComplex(os.Getenv("CHANNEL_ID"), &message)
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(j.ID())

	s.Start()

	select {
	case <-time.After(time.Minute):
	}

	err = s.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}

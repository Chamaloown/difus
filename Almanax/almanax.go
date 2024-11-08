package almanax

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	db "github.com/chamaloown/difus/Almanax/Db"
	gocron "github.com/go-co-op/gocron/v2"
)

func formatAlmanax(almanax db.Almanax) string {
	return fmt.Sprintf(
		"📅 Almanax du **%s**\n\n🔮 **Méryde** : %s\n📈 **Type de Bonus** : %s\n🎁 **Bonus** : %s\n🎒 **Offrande** : %s x%d\n💰 **Prix estimé** : %d kamas\n",
		almanax.Date.Format("2006/01/02"), almanax.Merydes, almanax.Type, almanax.Bonus, almanax.Offerings, almanax.QuantityOffered, almanax.Kamas)
}

func formatWeeklyAlmanax(almanaxes []db.Almanax) string {
	var result string

	for _, almanax := range almanaxes {
		result += fmt.Sprintf(
			"📅 Almanax du **%s**\n🔮 **Méryde** : %s\n📈 **Type de Bonus** : %s\n🎁 **Bonus** : %s\n🎒 **Offrande** : %s x%d\n💰 **Prix estimé** : %d kamas\n\n",
			almanax.Date.Format("2006/01/02"), almanax.Merydes, almanax.Type, almanax.Bonus, almanax.Offerings, almanax.QuantityOffered, almanax.Kamas,
		)
	}
	return result
}


func GetAlmanax(message string) discordgo.MessageSend {
	datePattern := `.*(\b\d{2}/\d{2}/\d{4}\b).*`
	dbInstance := db.GetDBInstance()
	re := regexp.MustCompile(datePattern)

	switch {
		case re.MatchString(message):
			dateStr := re.FindStringSubmatch(message)
			date, err := time.Parse("02/01/2006", dateStr[1])
			if err != nil {
				log.Fatal(err)
			}
			almanax, err := db.GetAlmanax(dbInstance, date)
			if err != nil {
				log.Fatal(err)
			}
			return discordgo.MessageSend{
				Content: formatAlmanax(almanax),
			}

		case strings.Contains(message, "week"):
			almanaxes, err := db.GetWeeklyAlmanax(dbInstance)
			if err != nil {
				log.Fatal(err)
			}
			return discordgo.MessageSend{
				Content: formatWeeklyAlmanax(almanaxes),
			}
			
		default:
			almanax, err := db.GetAlmanax(dbInstance, time.Now())
			if err != nil {
				log.Fatal(err)
			}
			return discordgo.MessageSend{
				Content: formatAlmanax(almanax),
			}
	}
} 

func Run(discord *discordgo.Session) {
	fmt.Println("Loading almanax...")
	db.Setup()
	fmt.Println("Almanax Loaded!")

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.NewJob(
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
	s.Start()

	select {
	case <-time.After(time.Minute):
	}

	err = s.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}

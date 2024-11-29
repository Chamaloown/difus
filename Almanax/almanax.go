package almanax

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	formatter "github.com/chamaloown/difus/Almanax/Formatter"
	parser "github.com/chamaloown/difus/Almanax/Parser"
	reader "github.com/chamaloown/difus/Almanax/Reader"
	writer "github.com/chamaloown/difus/Almanax/Writer"
	database "github.com/chamaloown/difus/Database"
	models "github.com/chamaloown/difus/Models"
	gocron "github.com/go-co-op/gocron/v2"
)



func setup() {
	db := database.GetDBInstance()

	fmt.Println("Is almanax complet")
	if reader.IsAlmanaxComplet(db) {
		fmt.Println("Database is already set to use!")
		return
	}
	
	fmt.Println("parser Run")
	records, err := parser.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("records", records)

	for _, val := range records[1:] {
		date, err := time.Parse("02/01/2006", val[0])
		if err != nil {
			log.Fatal(err)
		}

		qty, err := strconv.Atoi(val[5])
		if err != nil {
			log.Fatal(err)
		}

		kamas, err := strconv.Atoi(val[6])
		if err != nil {
			log.Fatal(err)
		}

		newEntry := models.Almanax{
			Date:            date,
			Merydes:         val[1],
			Type:            val[2],
			Bonus:           val[3],
			Offerings:       val[4],
			QuantityOffered: qty,
			Kamas:           kamas,
		}

		_, err = writer.CreateAlmanax(db, newEntry)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Run(discord *discordgo.Session) {
	fmt.Println("Loading almanax...")
	setup()
	fmt.Println("Successfully charged the database!, Almanax loaded!")

	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(1, 30, 0),
				gocron.NewAtTime(21, 0, 0),
			),
		),
		gocron.NewTask(
			func() {
				pg := database.GetDBInstance()
				alamanax, err := reader.GetAlmanax(pg, time.Now().AddDate(1, 0, 0))
				if err != nil {
					log.Fatal(err)
				}
				var message = discordgo.MessageSend{
					Content: formatter.FormatAlmanax(alamanax),
				}
				discord.ChannelMessageSendComplex(os.Getenv("CHANNEL_ID"), &message)
			},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	s.Start()

	<-time.After(time.Minute)

	err = s.Shutdown()
	if err != nil {
		log.Fatal(err)
	}
}

func GetAlmanax(message string) discordgo.MessageSend {
	datePattern := `.*(\b\d{2}/\d{2}/\d{4}\b).*`
	dbInstance := database.GetDBInstance()
	re := regexp.MustCompile(datePattern)

	switch {
		case re.MatchString(message):
			dateStr := re.FindStringSubmatch(message)
			date, err := time.Parse("02/01/2006", dateStr[1])
			if err != nil {
				log.Fatal(err)
			}
			almanax, err := reader.GetAlmanax(dbInstance, date)
			if err != nil {
				log.Fatal(err)
			}
			return discordgo.MessageSend{
				Content: formatter.FormatAlmanax(almanax),
			}

		case strings.Contains(message, "week"):
			almanaxes, err := reader.GetWeeklyAlmanax(dbInstance)
			if err != nil {
				log.Fatal(err)
			}
			return discordgo.MessageSend{
				Content: formatter.FormatWeeklyAlmanax(almanaxes),
			}
			
		default:
			almanax, err := reader.GetAlmanax(dbInstance, time.Now())
			if err != nil {
				log.Fatal(err)
			}
			return discordgo.MessageSend{
				Content: formatter.FormatAlmanax(almanax),
			}
	}
}

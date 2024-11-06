package almanax

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)

func Run() {
	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		fmt.Println("Error!")
	}

	// add a job to the scheduler
	j, err := s.NewJob(
		gocron.DailyJob(
			1,
			gocron.NewAtTimes(
				gocron.NewAtTime(0, 30, 0),
				gocron.NewAtTime(22, 0, 0),
			),
		),
		gocron.NewTask(
			func(a string, b int) {
				fmt.Println("Almanax!")
			},
			"hello",
			1,
		),	
	)
	if err != nil {
		fmt.Println("Error!")

	}
	// each job has a unique id
	fmt.Println(j.ID())

	// start the scheduler
	s.Start()

	// block until you are ready to shut down
	select {
	case <-time.After(time.Minute):
	}

	// when you're done, shut it down
	err = s.Shutdown()
	if err != nil {
		fmt.Println("Error!")
	}
}
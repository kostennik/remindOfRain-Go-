package main

import (
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
	"remind-of-rain/src/app"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("0 20 * * *", func() {
		err := app.Start()
		if err != nil {
			log.Err(err).Msg("error while running application")
			return
		}
	})
	c.Start()

	for {
		time.Sleep(1 * time.Second)
	}
}

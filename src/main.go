package main

import (
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog/log"
	log2 "log"
	"remind-of-rain/src/app"
	"remind-of-rain/src/config"
	"time"
)

func main() {
	cfg, err := config.NewConfiguration("./config/properties.yaml").LoadConfig()
	if err != nil {
		log2.Fatal(err)
	}

	c := cron.New()
	c.AddFunc(cfg.App.EventTime, func() {
		err := app.Start(cfg)
		if err != nil {
			log.Err(err).Msg("error while running application")
		}
	})
	c.Start()

	for {
		time.Sleep(1 * time.Second)
	}
}

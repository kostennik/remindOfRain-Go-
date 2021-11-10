package app

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"remind-of-rain/src/config"
	"remind-of-rain/src/messenger"
	"remind-of-rain/src/weather"
)

func Start(cfg *config.Configuration) error {
	ctx := context.Background()

	w := weather.NewAccuweather(cfg.Weather.Accuweather.ApiKey, cfg.Weather.Accuweather.Url, cfg.Weather.Accuweather.CityCode, cfg.Weather.Accuweather.Language)
	forecast, err := w.GetForecast(ctx)
	if err != nil {
		return errors.Wrap(err, "error while getting weather")
	}

	detect, err := forecast.RainDetect(cfg.Weather.Accuweather.Language)
	if err != nil {
		return err
	}

	//if the rain is not detected, don't send the message
	if !detect {
		return nil
	}

	p := messenger.NewPushover(cfg.Messenger.Pushover.AppKey, cfg.Messenger.Pushover.UserKey)
	msg := fmt.Sprintf("Weather for tomorrow: %1.f-%1.f, day: %s, night: %s", forecast.TempMin, forecast.TempMax, forecast.DescriptionDay, forecast.DescriptionNight)
	err = p.SendMessage(ctx, "Take an umbrella", msg)
	if err != nil {
		return errors.Wrap(err, "error while sending message")
	}
	return nil
}

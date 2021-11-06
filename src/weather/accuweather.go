package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
	client "remind-of-rain/src/httpClient"
	"time"
)

type accuweather struct {
	ApiKey     string
	Url        string
	CityCode   string
	Language   string
	httpGetter client.HttpClient
}

func NewAccuweather(apiKey string, url string, cityCode string, language string) *accuweather {
	return &accuweather{
		ApiKey:     apiKey,
		Url:        url,
		CityCode:   cityCode,
		Language:   language,
		httpGetter: client.NewHttpClient(20 * time.Second),
	}
}

func (a accuweather) GetForecast(ctx context.Context) (*weather, error) {
	log.Debug().Msg("starting GetForecast()")
	url := fmt.Sprintf("%s/%s?apikey=%s&language=%s&details=%v&metric=%v", a.Url, a.CityCode, a.ApiKey, a.Language, true, true)

	resp, err := a.httpGetter.Do(ctx, url, http.MethodGet, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "GetForecast(): error while getting a forecast")
	}

	var forecast = new(accuweatherForecast)
	if err = json.Unmarshal(resp, forecast); err != nil {
		return nil, errors.Wrapf(err, "GetForecast(): error while encoding struct from json")
	}

	var result = new(weather)
	if forecast != nil || forecast.DailyForecast != nil {
		result = &weather{
			TempMin:          forecast.DailyForecast[1].Temperature.Minimum.Value,
			TempMax:          forecast.DailyForecast[1].Temperature.Maximum.Value,
			DescriptionDay:   forecast.DailyForecast[1].Day.LongPhrase,
			DescriptionNight: forecast.DailyForecast[1].Night.LongPhrase,
		}
	}

	log.Debug().Msg("end GetForecast()")
	return result, nil
}

type accuweatherForecast struct {
	DailyForecast []dailyForecast `json:"DailyForecasts"`
}

type dailyForecast struct {
	Temperature *temperature `json:"Temperature"`
	Day         *day         `json:"Day"`
	Night       *day         `json:"Night"`
}

type temperature struct {
	Minimum *unit `json:"Minimum"`
	Maximum *unit `json:"Maximum"`
}

type unit struct {
	Value    float64 `json:"Value"`
	Unit     string  `json:"Unit"`
	UnitType int     `json:"UnitType"`
}

type day struct {
	LongPhrase string `json:"LongPhrase"`
}

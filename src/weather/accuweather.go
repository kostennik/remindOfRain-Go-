package weather

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type accuweather struct {
	ApiKey   string
	Url      string
	CityCode string
	Language string
	httpGet  func(url string) (resp *http.Response, err error)
}

func NewAccuweather(apiKey string, url string, cityCode string, language string) *accuweather {
	return &accuweather{
		ApiKey:   apiKey,
		Url:      url,
		CityCode: cityCode,
		Language: language,
		httpGet:  http.Get,
	}
}

func (a accuweather) GetWeather() (*weather, error) {
	url := fmt.Sprintf("%s/%s?apikey=%s&language=%s&details=%v&metric=%v", a.Url, a.CityCode, a.ApiKey, a.Language, true, true)

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting the data from: %s", a.Url)
	}
	defer resp.Body.Close()

	responseBodyRaw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error occurred while reading response body")
	}

	var result = new(accuweatherForecast)
	if err = json.Unmarshal(responseBodyRaw, result); err != nil {
		return nil, errors.Wrapf(err, "error while encoding struct from json")
	}

	if result == nil {
		return nil, errors.New("accuweather forecast is empty")
	}

	return &weather{
		TempMin:          result.DailyForecast[1].Temperature.Minimum.Value,
		TempMax:          result.DailyForecast[1].Temperature.Maximum.Value,
		DescriptionDay:   result.DailyForecast[1].Day.LongPhrase,
		DescriptionNight: result.DailyForecast[1].Night.LongPhrase,
	}, nil

}

type accuweatherForecast struct {
	DailyForecast []dailyForecast `json:"DailyForecasts"`
}

type dailyForecast struct {
	Temperature *temperature `json:"Temperature"`
	Day         *day         `json:"Day"`
	Night       *night       `json:"Night"`
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
type night struct {
	LongPhrase string `json:"LongPhrase"`
}

package weather

import (
	"context"
	"fmt"
	"strings"
)

const (
	engLang = "en-US"
	plLang  = "pl-pl"
)

var (
	engPhrases = []string{
		"rain", "storm", "precip", "deluge", "showers",
	}
	plPhrases = []string{
		"deszcz", "opad", "grad", "ulewa", "burza",
	}
)

type Weather interface {
	GetForecast(ctx context.Context) (*weather, error)
}

type weather struct {
	language         string
	TempMin          float64
	TempMax          float64
	DescriptionDay   string
	DescriptionNight string
}

func (w weather) RainDetect(language string) (bool, error) {
	switch language {
	case engLang:
		return w.contains(engPhrases), nil
	case plLang:
		return w.contains(plPhrases), nil
	default:
		return false, errLanguageDoesNotSupport
	}
}

func (w weather) contains(phrases []string) bool {
	for _, p := range phrases {
		if strings.Contains(fmt.Sprintf("%s %s", w.DescriptionDay, w.DescriptionNight), p) {
			return true
		}
	}
	return false
}

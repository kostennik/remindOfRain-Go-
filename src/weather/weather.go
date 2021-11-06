package weather

import "context"

type Weather interface {
	GetForecast(ctx context.Context) (*weather, error)
}

type weather struct {
	TempMin          float64
	TempMax          float64
	DescriptionDay   string
	DescriptionNight string
}

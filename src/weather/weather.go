package weather

type Weather interface {
	GetForecast() (*weather, error)
}

type weather struct {
	TempMin          float64
	TempMax          float64
	DescriptionDay   string
	DescriptionNight string
}

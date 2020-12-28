package weather

type Weather interface {
	GetWeather() (*weather, error)
}

type weather struct {
	TempMin          float64
	TempMax          float64
	DescriptionDay   string
	DescriptionNight string
}

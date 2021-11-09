package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

func NewConfiguration(path string) *Configuration {
	return &Configuration{path: path}
}

type Configuration struct {
	path      string
	App       *app       `yaml:"app"`
	Weather   *weather   `yaml:"weather"`
	Messenger *messenger `yaml:"messenger"`
}

type app struct {
	EventTime string `yaml:"eventTime"`
}

type weather struct {
	Accuweather *accuweather `yaml:"accuweather"`
}

type messenger struct {
	Pushover *pushover `yaml:"pushover"`
}

type accuweather struct {
	Url      string `yaml:"url"`
	CityCode int    `yaml:"cityCode"`
	ApiKey   string `yaml:"apiKey"`
	Language string `yaml:"language"`
}

type pushover struct {
	Url     string `yaml:"url"`
	AppKey  string `yaml:"appKey"`
	UserKey string `yaml:"userKey"`
}

func (c Configuration) LoadConfig() (*Configuration, error) {
	file, err := ioutil.ReadFile(c.path)
	if err != nil {
		return nil, err
	}

	if len(file) == 0 {
		return nil, errFilePropertyIsEmpty
	}

	var conf Configuration

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

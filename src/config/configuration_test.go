package config

import (
	"os"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type testFile struct {
		name string
		body []byte
	}

	tests := []struct {
		name       string
		testConfig testFile
		want       *configuration
		wantErr    bool
	}{
		{
			name: "correct read configuration from yaml-file",
			testConfig: testFile{
				name: "./correct-config.yaml",
				body: []byte(`
weather:
  accuweather:
    url: https://dataservice.accuweather.com/forecasts/v1/daily/5day
    cityCode: 434334 #see in https://developer.accuweather.com/
    apiKey: rfrfrggdsdfsgf
    language: en-US
messenger:
  pushover:
    url: https://api.pushover.net/1/messages.json
    appKey: tyhgeedrtgreg
    userKey: kliykjul786uiju
`),
			},
			want: &configuration{
				Weather: &weather{
					Accuweather: &accuweather{
						Url:      "https://dataservice.accuweather.com/forecasts/v1/daily/5day",
						CityCode: 434334,
						ApiKey:   "rfrfrggdsdfsgf",
						Language: "en-US",
					},
				},
				Messenger: &messenger{
					Pushover: &pushover{
						Url:     "https://api.pushover.net/1/messages.json",
						AppKey:  "tyhgeedrtgreg",
						UserKey: "kliykjul786uiju",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "try to read from empty yaml-file",
			testConfig: testFile{
				name: "./empty-config.yaml",
				body: []byte(``),
			},
			wantErr: true,
		},
		{
			name: "try to read from missing yaml-file",
			testConfig: testFile{
				name: "./missing-config.yaml",
				body: []byte(``),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := os.WriteFile(tt.testConfig.name, tt.testConfig.body, 0777)
			if err != nil {
				t.Errorf("error while creating new test config file")
				return
			}

			defer func() {
				os.Remove(tt.testConfig.name)
			}()

			got, err := NewConfiguration(tt.testConfig.name).LoadConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func createTestFile(name string, body []byte) error {
	err := os.WriteFile(name, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

package weather

import (
	"context"
	"errors"
	"io"
	"reflect"
	"remind-of-rain/src/httpClient"
	"testing"
)

func Test_accuweather_GetWeather(t *testing.T) {
	type fields struct {
		ApiKey     string
		httpGetter httpClient.HttpClient
	}

	tests := []struct {
		name    string
		fields  fields
		want    *weather
		wantErr bool
	}{
		{
			name: "test correct get weather",
			fields: fields{
				ApiKey: "correct-key",
				httpGetter: httpGetterMock{
					t: t,
					DoMock: func(ctx context.Context, url, method string, body io.Reader) ([]byte, error) {
						return correctResponse, nil
					},
				},
			},
			want: &weather{
				TempMin:          6,
				TempMax:          8.2,
				DescriptionDay:   "Breezy this morning; cloudy with a brief shower or two",
				DescriptionNight: "Mostly cloudy with a shower in places",
			},
			wantErr: false,
		},
		{
			name: "try to get weather with incorrect api-key",
			fields: fields{
				ApiKey: "incorrect-key",
				httpGetter: httpGetterMock{
					t: t,
					DoMock: func(ctx context.Context, url, method string, body io.Reader) ([]byte, error) {
						return unauthorizedResponse, errors.New("authorization failed")
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := accuweather{
				ApiKey:     tt.fields.ApiKey,
				httpGetter: tt.fields.httpGetter,
			}
			got, err := a.GetForecast(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetForecast() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetForecast() got = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}

type httpGetterMock struct {
	t      *testing.T
	DoMock func(ctx context.Context, url, method string, body io.Reader) ([]byte, error)
}

func (h httpGetterMock) Do(ctx context.Context, url, method string, body io.Reader) ([]byte, error) {
	if url == "" || method == "" {
		h.t.Error("Do(): url and method is required")
	}
	return h.DoMock(ctx, url, method, body)
}

var unauthorizedResponse = []byte(`
{
  "Code": "Unauthorized",
  "Message": "Api Authorization failed",
  "Reference": "/forecasts/v1/daily/5day/270933?apikey=incorrect-key&language=en-US&details=true&metric=true"
}
`)

var correctResponse = []byte(`
{
  "Headline": {
    "EffectiveDate": "2021-09-05T19:00:00+01:00",
    "EffectiveEpochDate": 1636135200,
    "Severity": 5,
    "Text": "Expect showers Friday night",
    "Category": "rain",
    "EndDate": "2021-09-06T07:00:00+01:00",
    "EndEpochDate": 1636178400,
    "MobileLink": "http://www.accuweather.com/en/pl/stawiguda/270933/daily-weather-forecast/270933?unit=c&lang=en-us",
    "Link": "http://www.accuweather.com/en/pl/stawiguda/270933/daily-weather-forecast/270933?unit=c&lang=en-us"
  },
  "DailyForecasts": [
	{},
    {
      "Date": "2021-09-05T07:00:00+01:00",
      "EpochDate": 1636092000,
      "Sun": {
        "Rise": "2021-09-05T06:45:00+01:00",
        "EpochRise": 1636091100,
        "Set": "2021-09-05T15:58:00+01:00",
        "EpochSet": 1636124280
      },
      "Moon": {
        "Rise": "2021-09-05T07:18:00+01:00",
        "EpochRise": 1636093080,
        "Set": "2021-09-05T16:19:00+01:00",
        "EpochSet": 1636125540,
        "Phase": "WaxingCrescent",
        "Age": 1
      },
      "Temperature": {
        "Minimum": {
          "Value": 6,
          "Unit": "C",
          "UnitType": 17
        },
        "Maximum": {
          "Value": 8.2,
          "Unit": "C",
          "UnitType": 17
        }
      },
      "RealFeelTemperature": {
        "Minimum": {
          "Value": 1.4,
          "Unit": "C",
          "UnitType": 17
        },
        "Maximum": {
          "Value": 5.5,
          "Unit": "C",
          "UnitType": 17
        }
      },
      "RealFeelTemperatureShade": {
        "Minimum": {
          "Value": 1.4,
          "Unit": "C",
          "UnitType": 17
        },
        "Maximum": {
          "Value": 5.5,
          "Unit": "C",
          "UnitType": 17
        }
      },
      "HoursOfSun": 1.2,
      "DegreeDaySummary": {
        "Heating": {
          "Value": 11,
          "Unit": "C",
          "UnitType": 17
        },
        "Cooling": {
          "Value": 0,
          "Unit": "C",
          "UnitType": 17
        }
      },
      "AirAndPollen": [
        {
          "Name": "AirQuality",
          "Value": 0,
          "Category": "Good",
          "CategoryValue": 1,
          "Type": "Ozone"
        },
        {
          "Name": "Grass",
          "Value": 0,
          "Category": "Low",
          "CategoryValue": 1
        },
        {
          "Name": "Mold",
          "Value": 0,
          "Category": "Low",
          "CategoryValue": 1
        },
        {
          "Name": "Ragweed",
          "Value": 0,
          "Category": "Low",
          "CategoryValue": 1
        },
        {
          "Name": "Tree",
          "Value": 0,
          "Category": "Low",
          "CategoryValue": 1
        },
        {
          "Name": "UVIndex",
          "Value": 0,
          "Category": "Low",
          "CategoryValue": 1
        }
      ],
      "Day": {
        "Icon": 12,
        "IconPhrase": "Showers",
        "HasPrecipitation": true,
        "PrecipitationType": "Rain",
        "PrecipitationIntensity": "Light",
        "ShortPhrase": "A brief shower or two",
        "LongPhrase": "Breezy this morning; cloudy with a brief shower or two",
        "PrecipitationProbability": 70,
        "ThunderstormProbability": 14,
        "RainProbability": 69,
        "SnowProbability": 0,
        "IceProbability": 0,
        "Wind": {
          "Speed": {
            "Value": 20.4,
            "Unit": "km/h",
            "UnitType": 7
          },
          "Direction": {
            "Degrees": 224,
            "Localized": "SW",
            "English": "SW"
          }
        },
        "WindGust": {
          "Speed": {
            "Value": 35.2,
            "Unit": "km/h",
            "UnitType": 7
          },
          "Direction": {
            "Degrees": 210,
            "Localized": "SSW",
            "English": "SSW"
          }
        },
        "TotalLiquid": {
          "Value": 3.8,
          "Unit": "mm",
          "UnitType": 3
        },
        "Rain": {
          "Value": 3.8,
          "Unit": "mm",
          "UnitType": 3
        },
        "Snow": {
          "Value": 0,
          "Unit": "cm",
          "UnitType": 4
        },
        "Ice": {
          "Value": 0,
          "Unit": "mm",
          "UnitType": 3
        },
        "HoursOfPrecipitation": 2.5,
        "HoursOfRain": 2.5,
        "HoursOfSnow": 0,
        "HoursOfIce": 0,
        "CloudCover": 98,
        "Evapotranspiration": {
          "Value": 0,
          "Unit": "mm",
          "UnitType": 3
        },
        "SolarIrradiance": {
          "Value": 0.3,
          "Unit": "W/m²",
          "UnitType": 33
        }
      },
      "Night": {
        "Icon": 40,
        "IconPhrase": "Mostly cloudy w/ showers",
        "HasPrecipitation": true,
        "PrecipitationType": "Rain",
        "PrecipitationIntensity": "Light",
        "ShortPhrase": "Mostly cloudy with a shower",
        "LongPhrase": "Mostly cloudy with a shower in places",
        "PrecipitationProbability": 40,
        "ThunderstormProbability": 8,
        "RainProbability": 40,
        "SnowProbability": 0,
        "IceProbability": 0,
        "Wind": {
          "Speed": {
            "Value": 14.8,
            "Unit": "km/h",
            "UnitType": 7
          },
          "Direction": {
            "Degrees": 244,
            "Localized": "WSW",
            "English": "WSW"
          }
        },
        "WindGust": {
          "Speed": {
            "Value": 24.1,
            "Unit": "km/h",
            "UnitType": 7
          },
          "Direction": {
            "Degrees": 247,
            "Localized": "WSW",
            "English": "WSW"
          }
        },
        "TotalLiquid": {
          "Value": 0.5,
          "Unit": "mm",
          "UnitType": 3
        },
        "Rain": {
          "Value": 0.5,
          "Unit": "mm",
          "UnitType": 3
        },
        "Snow": {
          "Value": 0,
          "Unit": "cm",
          "UnitType": 4
        },
        "Ice": {
          "Value": 0,
          "Unit": "mm",
          "UnitType": 3
        },
        "HoursOfPrecipitation": 1,
        "HoursOfRain": 1,
        "HoursOfSnow": 0,
        "HoursOfIce": 0,
        "CloudCover": 91,
        "Evapotranspiration": {
          "Value": 0,
          "Unit": "mm",
          "UnitType": 3
        },
        "SolarIrradiance": {
          "Value": 0,
          "Unit": "W/m²",
          "UnitType": 33
        }
      },
      "Sources": [
        "AccuWeather"
      ],
      "MobileLink": "http://www.accuweather.com/en/pl/stawiguda/270933/daily-weather-forecast/270933?day=1&unit=c&lang=en-us",
      "Link": "http://www.accuweather.com/en/pl/stawiguda/270933/daily-weather-forecast/270933?day=1&unit=c&lang=en-us"
    }
  ]
}
`)

package weather

import (
	"errors"
	"testing"
)

func Test_weather_RainCheck(t *testing.T) {
	type fields struct {
		language         string
		TempMin          float64
		TempMax          float64
		DescriptionDay   string
		DescriptionNight string
	}

	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr error
	}{
		{
			name: "checking for English phrase",
			fields: fields{
				language:         "en-US",
				DescriptionDay:   "Sun",
				DescriptionNight: "Partly cloudy with a few showers",
			},
			want: true,
		},
		{
			name: "checking for Poland phrase",
			fields: fields{
				language:         "pl-pl",
				DescriptionDay:   "Zachmurzenie duże z przelotnymi opadami",
				DescriptionNight: "jasna pogoda",
			},
			want: true,
		},
		{
			name: "try to checking for missing phrase",
			fields: fields{
				language:         "en-US",
				DescriptionDay:   "Periods of clouds and sun",
				DescriptionNight: "-",
			},
			want: false,
		},
		{
			name: "try to check with unsupported language",
			fields: fields{
				language:         "de-DE",
				DescriptionDay:   "Teils bewölkt mit einigen Schauern",
				DescriptionNight: "klares Wetter",
			},
			wantErr: errLanguageDoesNotSupport,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := weather{
				language:         tt.fields.language,
				TempMin:          tt.fields.TempMin,
				TempMax:          tt.fields.TempMax,
				DescriptionDay:   tt.fields.DescriptionDay,
				DescriptionNight: tt.fields.DescriptionNight,
			}
			got, err := w.RainDetect(tt.fields.language)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("RainDetect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RainDetect() got = %v, want %v", got, tt.want)
			}
		})
	}
}

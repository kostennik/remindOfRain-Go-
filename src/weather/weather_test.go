package weather

import "testing"

func TestValidate(t *testing.T) {
	type test struct {
		name string
		wanted Weather
		pass   bool
	}

	tests := []test{
		{
			name: "Test 1. Pogoda sloneczna",
			wanted: Weather{
				TempMin:          "3",
				TempMax:          "13",
				DescriptionDay:   "Będzie slonce",
				DescriptionNight: "Będzie zimno",
			},
			pass: false,
		},
		{
			name: "Test 2. Opady",
			wanted: Weather{
				TempMin:          "5",
				TempMax:          "8",
				DescriptionDay:   "Oczekują się opady",
				DescriptionNight: "",
			},
			pass: true,
		},
		{
			name: "Test 3. Deszcz",
			wanted: Weather{
				TempMin:          "16",
				TempMax:          "25",
				DescriptionDay:   "Oczekuje się deszcz, duże opady",
				DescriptionNight: "",
			},
			pass:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := Validate(tt.wanted)
			if status != tt.pass {
				t.Errorf("error while validate data")
			}
		})

	}
}

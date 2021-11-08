package app

import "testing"

func TestStart(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "correct run application",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Start(); (err != nil) != tt.wantErr {
				t.Errorf("Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

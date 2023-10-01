package fingerprint

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewFingerprintBuilder(t *testing.T) {
	tests := []struct {
		name    string
		want    *Builder
		wantErr bool
	}{
		{
			name: "gen_json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFingerprintBuilder("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
			if err != nil {
				panic(err)
			}

			got.GenerateProfile()

			n, err := got.Build()
			if err != nil {
				panic(err)
			}

			out, err := json.Marshal(n)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(out))
		})
	}
}

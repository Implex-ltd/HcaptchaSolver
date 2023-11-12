package fingerprint

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	JWT_TOKEN = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJmIjowLCJzIjoyLCJ0IjoidyIsImQiOiIrOVZwZU5nUUFwSUpXR2dhY0Y1LzlhTHhad0l5ZHpoQzVwMUw4QWh1cXRDTTJIQjlyOHpJeFNWczdGWW9HOU1uVW54U1I4aEkxanFoUG1YRE5pdGFaK2NDOUM0UWN2NlRTZUhISzI4Nzc5T0pTRVIyOUhDOTdUNysvVUFvQjVReXkzYi9XTkpjMnBUbk1hTDA1RkNYV2QzMmMwa1krdFYxVUlsMzdEekJpdUVQc1kzNDdPSHRYdmpXRHc9PVcrQTlSRTBnNlB4OGdVRlciLCJsIjoiaHR0cHM6Ly9uZXdhc3NldHMuaGNhcHRjaGEuY29tL2MvNzhlZTZmYyIsImkiOiJzaGEyNTYtdEs3YTVnbXE3Wjd1R0w2REh5OW9ReHUvRmsvdW1WdzNlTFBaWitlS2lkdz0iLCJlIjoxNjk3MjQxNzE3LCJuIjoiaHN3IiwiYyI6MTAwMH0.AbJ3rc1gTbM0ADacD-MdrcCyrShmMmhYi98juOCT1aEdTaZ5z5AQcNZIfMuMkxG3gAbsgrqeyK_OiTHeoNIw5hLXrDeDGihgHVAF2A0gBbNpWVI78rACHJQn4UyP1GHoYTTu3-HeAQXmVP0JEd6X7UgAWUeIn0IRKLlrNrmO0aI"
)

func TestNewFingerprintBuilder(t *testing.T) {
	type args struct {
		useragent string
	}
	tests := []struct {
		name    string
		args    args
		want    *Builder
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				useragent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36 Edg/117.0.2045.60",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFingerprintBuilder(tt.args.useragent, "https://discord.com")
			if err != nil {
				panic(err)
			}

			if _, err := got.GenerateProfile(); err != nil {
				panic(err)
			}

			n, err := got.Build(JWT_TOKEN, false, true)
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

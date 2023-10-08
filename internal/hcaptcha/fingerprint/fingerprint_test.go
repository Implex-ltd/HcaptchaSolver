package fingerprint

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	jwtToken = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJmIjowLCJzIjoyLCJ0IjoidyIsImQiOiJLckwvMUtwdCszMmliMi9xdHlxV2orZzZxTlhsUFdETXVmY2xlbWV0RklJbVJFNGVDQ0tOOWpQQ1FEeTYwZThTejd0c2NoVlJ6V0xzcjJrYS9Ca1JDTW13dmVQcFM2YVZBLy9mN2l3QWxqeFZqeE9FYzFXVWFoeXVUMnZ1dU0wdXV4OWhuaVUrVlplM1grV0VFWUhqK3gvU05YVVV6aVpBUzBYYmdSOHFPTW1QUERCWkN0dXFGWjRBT0E9PXU4elZxbHRibU9lYUNyQ2wiLCJsIjoiaHR0cHM6Ly9uZXdhc3NldHMuaGNhcHRjaGEuY29tL2MvN2E3ZmMzZCIsImkiOiJzaGEyNTYtOWJZYUQxSGhUUG5EWURLWE52Q0ZZMFJ1NDVSdEE5dUtFd2RSYlVkNGc0MD0iLCJlIjoxNjk2NzQ1Nzk4LCJuIjoiaHN3IiwiYyI6MTAwMH0.hlkPmTRWICCp3vFP4eosMIxGp5RfIez-s7zAoqaLYjSgIKj0oBYLtZYUkNsZL1OY5ybxBtvWIqX1sbifQmjZPwQvgaeo1R9H7HZFxQKSGjDQAMGzP8WI0vlRwil-BjcQDrTjomghT-Fi0R4_tnuMv5DkMrTExpkBBz-qEjFDjxg"
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

			n, err := got.Build(jwtToken)
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

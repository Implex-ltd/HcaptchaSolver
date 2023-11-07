/**
  https://crates.io/crates/rust-hashcash/0.3.3
  Thanks hcaptcha 1990 algo !!

  fork: github.com/catalinc/hashcash
*/

package fingerprint

import (
	"fmt"
	"testing"
)

func TestGetStamp(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				data: "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJmIjowLCJzIjoyLCJ0IjoidyIsImQiOiJrcWU0Rk5kSmxqaFVsRVFMUVFoL3d2YlMrQWN2UC92d040OFN2ei9mRkQ4LzhGN0VwNHF4Mm5wRWpRNStWUE1CdVVQNFVmNFZBTVJGNVZWZzZRTjlMeHN3VXdoTzE5V0tHUDIxSTNKSGVFU1hYdTkwTlNmZmlZNFBuUU1mVmw0dG92MDN4SVk3aWtuQWFDUHdYMm5SeVVWMURrSzhsVzJ0djhPZzJqdjFFZHJQNmc0a2lYQUdxMm5pZ3c9PTVBaUxZZmZwREEwRFY3QXYiLCJsIjoiaHR0cHM6Ly9uZXdhc3NldHMuaGNhcHRjaGEuY29tL2MvNzhlZTZmYyIsImkiOiJzaGEyNTYtdEs3YTVnbXE3Wjd1R0w2REh5OW9ReHUvRmsvdW1WdzNlTFBaWitlS2lkdz0iLCJlIjoxNjk3NTc2MTk2LCJuIjoiaHN3IiwiYyI6MTAwMH0.n0lbRXee0rrYEiODWLmPneTi-lDMPZ6bGX7ar0lp5PzHfX6d_bIeV39WYnsv2V5aoDybj62KKzjIoUJJAkVbiGSmP-ngTgD8tb6PlAdzt-W48b1GMJbkkjXEP-UB4k8j9pmxSLgNegupvJi1XKfmnHh-R4k26cdz3yAUPQ19AD0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := ParseJWT(tt.args.data)
			if err != nil {
				panic(err)
			}

			for i := 0; i < 10; i++ {
				got, err := GetStamp(uint(token.Difficuly), token.PowData)
				if err != nil {
					panic(err)
				}

				fmt.Println(got)
			}
		})
	}
}

package utils

import (
	"fmt"
	"testing"
)

func TestLookup(t *testing.T) {
	type args struct {
		Address string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				Address: "158.46.169.117",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lookup(tt.args.Address)
			if err != nil {
				panic(err)
			}

			fmt.Println(got)
		})
	}
}

package events

import (
	"fmt"
	"testing"
)

func TestDecStr(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				input: EncStr("Google Inc. (NVIDIA)"),
			},
		},
		{
			name: "1",
			args: args{
				input: []string{"o3IzMxCzM1yjNWUD", "8", "b", "EKSYKCEHOCWTT"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.args.input)

			got := DecStr(tt.args.input)
			fmt.Println(got)
		})
	}
}

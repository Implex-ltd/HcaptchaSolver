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
				input: EncStr("Europe/Paris"),
			},
		},
		{
			name: "2",
			args: args{
				input: []string{"Atm2EDnXEjn5IZm4", "2", "d", "PMXWVNTKZWOZB"},
			},
		}, {
			name: "3",
			args: args{
				input: []string{"ZmyudN2AdmWgTOxq", "3", "2", "MZCWFVLBUQBQD"},
			},
		},
		{
			name: "3",
			args: args{
				input: []string{"M0ADnxQzM1Uzn4Et", "2", "b", "KNTRWSGCLYFRB"},
			},
		},
		{
			name: "3",
			args: args{
				input: []string{"eTM5iTj1iTjxiTjkhJMlyGBVFGD2qdrzUSMuUzMweJM4ezMwAdnljkMl", "7", "31", "BVQJYNQIERDPC"},
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

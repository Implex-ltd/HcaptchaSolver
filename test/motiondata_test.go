package test

import (
	"fmt"
	"testing"

	h "github.com/Implex-ltd/hcsolver/internal/hcaptcha"
)

func TestHcap_NewMotionData(t *testing.T) {
	resp, err := h.NewHcaptcha(&h.Config{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		SiteKey:   "requestBody.SiteKey",
		Domain:    "discord.com",
		Proxy:     "",
	})

	if err != nil {
		panic(err)
	}

	type args struct {
		m *h.Motion
	}
	tests := []struct {
		name string
		c    *h.Hcap
		args args
		want string
	}{
		{
			name: "lol",
			c:    resp,
			args: args{
				m: &h.Motion{
					IsCheck: false,
					Answers: map[string]string{"x": "true", "y": "true", "z": "true"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ot := tt.c.NewMotionData(tt.args.m)
			fmt.Println(ot)
		})
	}
}

package validator

type Settings struct {
	MinSubmitTime int
	MaxSubmitTime int
	Domain        string
	AlwaysText    bool
}

var (
	WebsiteSettings = map[string]Settings{
		"discord.com": {
			MinSubmitTime: 3200,
			MaxSubmitTime: 13000,
			AlwaysText:    true,
			Domain:        "discord.com",
		},
	}
)

func Validate(domain string) (*Settings, error) {
	settings, exists := WebsiteSettings[domain]
	if exists {
		return &settings, nil
	}

	return nil, nil
}

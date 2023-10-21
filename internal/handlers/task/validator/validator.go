package validator

type Settings struct {
	MinSubmitTime int
	MaxSubmitTime int
	Domain        string
	AlwaysText    bool
	OneclickOnly  bool
	Enabled       bool
	Rate          float64
}

var (
	WebsiteSettings = map[string]Settings{
		// friend req
		"a9b5fb07-92ff-493f-86fe-352a2803b3df": {
			MinSubmitTime: 3200,
			MaxSubmitTime: 13000,
			AlwaysText:    true,
			Domain:        "discord.com",
			Enabled:       true,
		},
		// join guild
		"b2b02ab5-7dae-4d6f-830e-7b55634c888b": {
			MinSubmitTime: 3200,
			MaxSubmitTime: 13000,
			AlwaysText:    true,
			Domain:        "discord.com",
			Enabled:       true,
		},
		// phone / email / login
		"f5561ba9-8f1e-40ca-9b5b-a0b3f719ef34": {
			MinSubmitTime: 3200,
			MaxSubmitTime: 13000,
			AlwaysText:    true,
			Domain:        "discord.com",
			Enabled:       true,
		},
		// register
		"4c672d35-0701-42b2-88c3-78380b0db560": {
			MinSubmitTime: 3200,
			MaxSubmitTime: 13000,
			AlwaysText:    true,
			Domain:        "discord.com",
			Enabled:       false,
		},
	}
)

func Validate(sitekey string) (*Settings, error) {
	settings, exists := WebsiteSettings[sitekey]
	if exists {
		return &settings, nil
	}

	return nil, nil
}

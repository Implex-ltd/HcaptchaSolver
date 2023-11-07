package user

type BodyAddBalance struct {
	Amount int    `json:"amount"`
	User   string `json:"user"`
}

type BodySetBypass struct {
	Enabled bool   `json:"enabled"`
	User    string `json:"user"`
}

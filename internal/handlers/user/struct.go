package user

type BodyAddBalance struct {
	Amount int    `json:"amount"`
	User   string `json:"user"`
}

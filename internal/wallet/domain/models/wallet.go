package models

type Wallet struct {
	UserId  string  `json:"userId"`
	Balance float64 `json:"balance"`
}

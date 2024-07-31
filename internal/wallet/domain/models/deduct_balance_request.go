package models

type DeductBalanceRequest struct {
	UserId string  `json:"userId"`
	Amount float64 `json:"amount"`
}

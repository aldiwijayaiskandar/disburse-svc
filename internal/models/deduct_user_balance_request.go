package models

type DeductUserBalanceRequest struct {
	UserId string `json:"userId"`
	Amount float64 `json:"amount"`
}
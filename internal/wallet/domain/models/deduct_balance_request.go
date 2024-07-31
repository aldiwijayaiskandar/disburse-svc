package models

type DeductBalanceRequest struct {
	UserId string  `json:"userId" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

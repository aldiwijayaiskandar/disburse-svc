package models

type DisburseRequest struct {
	UserId string  `json:"userId" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

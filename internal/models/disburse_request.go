package models

type DisburseRequest struct {
	UserId string `json:"userId"`
	Amount float64 `json:"amount"`
}
package models

type DisburseRequest struct {
	UserId string `json:"userId"`
	Amount string `json:"amount"`
}
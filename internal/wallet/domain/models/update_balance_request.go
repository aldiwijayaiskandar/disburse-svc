package models

type UpdateBalanceRequest struct {
	UserId  string  `json:"userId"`
	Balance float64 `json:"balance"`
}

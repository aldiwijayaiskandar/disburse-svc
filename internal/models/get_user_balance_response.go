package models

type GetUserBalanceResponse struct {
	Balance *float64 `json:"balance"`
	Error *string
}
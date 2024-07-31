package models

type GetUserBalanceRequest struct {
	UserId string `json:"userId" validate:"required"`
}

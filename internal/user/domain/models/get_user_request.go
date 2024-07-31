package models

type GetUserRequest struct {
	Id string `json:"id" validate:"required"`
}

package models

type GetUserResponse struct {
	Data *User `json:"data"`
	Error *string `json:"error"`
}
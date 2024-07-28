package models

type GetUserResponse struct {
	User *User `json:"user"`
	Error *string `json:"error"`
}
package models

type ApiResponse struct {
	Status  int     `json:"status"`
	Data    any     `json:"data"`
	Message *string `json:"message"`
}

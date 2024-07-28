package models

type ApiResponse struct {
	Status int32
	Data *interface{}
	Error *string
}
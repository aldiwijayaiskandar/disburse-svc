package models

type ApiResponse struct {
	Status int32
	Data any
	Error *string
}
package models

import (
	constants "github.com/paper-assessment/pkg/contants"
)

type GetUserResponse struct {
	Status    constants.StatusCode `json:"status"`
	User      *User                `json:"user"`
	ErrorCode constants.ErrorCode  `json:"errorCode"`
	Message   *string              `json:"message"`
}

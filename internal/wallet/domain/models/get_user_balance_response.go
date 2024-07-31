package models

import (
	constants "github.com/paper-assessment/pkg/contants"
)

type GetUserBalanceResponse struct {
	Status    constants.StatusCode `json:"status"`
	Balance   *float64             `json:"balance"`
	ErrorCode constants.ErrorCode  `json:"errorCode"`
	Message   *string              `json:"message"`
}

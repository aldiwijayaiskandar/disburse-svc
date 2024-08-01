package usecase

import (
	"github.com/paper-assessment/internal/disburse/data/services"
	"github.com/paper-assessment/internal/disburse/domain/models"
)

type DisburseUsecase struct {
	userService   services.UserServiceInterface
	walletService services.WalletServiceInterface
}

type DisburseUsecaseInterface interface {
	Disburse(request models.DisburseRequest, correlationId string) *models.DisburseResponse
}

func NewDisburseUsecase(userService services.UserServiceInterface, walletService services.WalletServiceInterface) DisburseUsecaseInterface {
	return &DisburseUsecase{
		userService:   userService,
		walletService: walletService,
	}
}

package mocks

import (
	"github.com/paper-assessment/internal/wallet/domain/models"
	"github.com/stretchr/testify/mock"
)

type WalletUsecase struct {
	mock.Mock
}

func (m *WalletUsecase) GetUserBalance(userId string) *models.GetUserBalanceResponse {
	args := m.Called(userId)
	return args.Get(0).(*models.GetUserBalanceResponse)
}

func (m *WalletUsecase) DeductBalance(request models.DeductBalanceRequest) *models.DeductBalanceResponse {
	args := m.Called(request)

	return args.Get(0).(*models.DeductBalanceResponse)
}

func (m *WalletUsecase) Reset() {
	m.ExpectedCalls = make([]*mock.Call, 0)
}

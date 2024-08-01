package mocks

import (
	"github.com/paper-assessment/internal/wallet/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockWalletService struct {
	mock.Mock
}

func (m *MockWalletService) GetUserBalance(request *models.GetUserBalanceRequest, correlationId string) models.GetUserBalanceResponse {
	args := m.Called(request, correlationId)

	return args.Get(0).(models.GetUserBalanceResponse)
}

func (m *MockWalletService) DeductUserBalance(request *models.DeductBalanceRequest, correlationId string) models.DeductBalanceResponse {
	args := m.Called(request, correlationId)

	return args.Get(0).(models.DeductBalanceResponse)
}

func (m *MockWalletService) Reset() {
	m.ExpectedCalls = make([]*mock.Call, 0)
}

package mocks

import (
	"context"

	"github.com/paper-assessment/internal/wallet/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) Get(ctx context.Context, userId string) (*models.Wallet, error) {
	args := m.Called(ctx, userId)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.Wallet), args.Error(1)
}

func (m *MockWalletRepository) DeductBalance(ctx context.Context, request models.DeductBalanceRequest) error {
	args := m.Called(ctx, request)

	if len(args) > 1 {
		return args.Error(1)
	}

	return nil
}

func (m *MockWalletRepository) Reset() {
	m.ExpectedCalls = make([]*mock.Call, 0)
}

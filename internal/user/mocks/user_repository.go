package mocks

import (
	"context"

	"github.com/paper-assessment/internal/user/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Get(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Reset() {
	m.ExpectedCalls = make([]*mock.Call, 0)
}

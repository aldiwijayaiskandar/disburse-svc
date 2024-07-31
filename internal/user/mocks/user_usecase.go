package mocks

import (
	"github.com/paper-assessment/internal/user/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) GetUser(request models.GetUserRequest) *models.GetUserResponse {
	args := m.Called(request)
	return args.Get(0).(*models.GetUserResponse)
}

func (m *MockUserUsecase) Reset() {
	m.ExpectedCalls = make([]*mock.Call, 0)
}

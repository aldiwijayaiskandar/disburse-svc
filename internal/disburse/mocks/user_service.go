package mocks

import (
	"github.com/paper-assessment/internal/user/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetUser(request *models.GetUserRequest, correlationId string) models.GetUserResponse {
	args := m.Called(request, correlationId)

	return args.Get(0).(models.GetUserResponse)
}

func (m *MockUserService) Reset() {
	m.ExpectedCalls = make([]*mock.Call, 0)
}

package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) Push(key string, replyTo string, body []byte, correlationId string) error {
	m.Called(key, replyTo, body, correlationId)

	return nil
}

func (m *MockPublisher) Reply(key string, body []byte, correlationId string) error {
	m.Called(key, body, correlationId)

	return nil
}

func (m *MockPublisher) ExpectedReplyAnyBody(key string, correlationId string) {
	m.On("Reply", key, mock.MatchedBy(func(b []byte) bool {
		return len(b) > 0
	}),
		correlationId).
		Return(nil)
}

func (m *MockPublisher) ExpectedPushAnyBody(key string, replyTo string, correlationId string) {
	m.On("Push", key, replyTo, mock.MatchedBy(func(b []byte) bool {
		return len(b) > 0
	}), correlationId).Return(nil)
}

func (m *MockPublisher) Reset() {
	m.ExpectedCalls = make([]*mock.Call, 0)
}

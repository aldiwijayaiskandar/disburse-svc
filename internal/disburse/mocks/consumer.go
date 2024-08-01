package mocks

import (
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/mock"
)

type MockConsumer struct {
	mock.Mock
}

func (m *MockConsumer) Listen(topics []string, listener func(delivery *amqp.Delivery)) error {
	args := m.Called(topics, listener)

	if args.Get(0) == nil {
		return nil
	}

	return args.Error(1)
}

func (m *MockConsumer) WaitReply(correlationId string) (*amqp.Delivery, error) {
	args := m.Called(correlationId)

	if args.Get(0) != nil {
		return args.Get(0).(*amqp.Delivery), args.Error(1)
	}

	return nil, args.Error(1)
}

func (m *MockConsumer) Reset() {
	m.ExpectedCalls = make([]*mock.Call, 0)
}

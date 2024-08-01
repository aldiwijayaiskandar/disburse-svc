package services

import (
	"encoding/json"

	"github.com/paper-assessment/internal/user/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/paper-assessment/pkg/rabbitmq"
)

type UserService struct {
	consumer  rabbitmq.ConsumerInterface
	publisher rabbitmq.PublisherInterface
}

type UserServiceInterface interface {
	GetUser(request *models.GetUserRequest, correlationId string) *models.GetUserResponse
}

func NewUserService(consumer rabbitmq.ConsumerInterface, publisher rabbitmq.PublisherInterface) UserServiceInterface {
	return &UserService{
		consumer:  consumer,
		publisher: publisher,
	}
}

func (s *UserService) GetUser(request *models.GetUserRequest, correlationId string) *models.GetUserResponse {
	// push to request user
	getUserRequestBody, _ := json.Marshal(&request)
	s.publisher.Push("user.get.request", getUserRequestBody, correlationId)

	// waiting for reply
	res, err := s.consumer.WaitReply(correlationId)

	if err != nil {
		// throw internal server error
		return &models.GetUserResponse{
			Status:    constants.Error,
			ErrorCode: constants.InternalServerError,
		}
	}

	var userRes models.GetUserResponse
	json.Unmarshal(res.Body, &userRes)

	return &userRes
}

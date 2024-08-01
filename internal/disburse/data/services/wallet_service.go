package services

import (
	"encoding/json"

	"github.com/paper-assessment/internal/wallet/domain/models"
	constants "github.com/paper-assessment/pkg/contants"
	"github.com/paper-assessment/pkg/rabbitmq"
)

type WalletService struct {
	consumer  rabbitmq.ConsumerInterface
	publisher rabbitmq.PublisherInterface
}

type WalletServiceInterface interface {
	GetUserBalance(request *models.GetUserBalanceRequest, correlationId string) *models.GetUserBalanceResponse
}

func NewWalletService(consumer rabbitmq.ConsumerInterface, publisher rabbitmq.PublisherInterface) WalletServiceInterface {
	return &WalletService{
		consumer:  consumer,
		publisher: publisher,
	}
}

func (s *WalletService) GetUserBalance(request *models.GetUserBalanceRequest, correlationId string) *models.GetUserBalanceResponse {
	// push to request user
	getUserRequestBody, _ := json.Marshal(&request)
	s.publisher.Push("wallet.balance.get.request", getUserRequestBody, correlationId)

	// waiting for reply
	res, err := s.consumer.WaitReply(correlationId)

	if err != nil {
		// throw internal server error
		return &models.GetUserBalanceResponse{
			Status:    constants.Error,
			ErrorCode: constants.InternalServerError,
		}
	}

	var walletRes models.GetUserBalanceResponse
	json.Unmarshal(res.Body, &walletRes)

	return &walletRes
}

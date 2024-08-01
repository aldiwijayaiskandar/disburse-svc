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
	GetUserBalance(request *models.GetUserBalanceRequest, correlationId string) models.GetUserBalanceResponse
	DeductUserBalance(request *models.DeductBalanceRequest, correlationId string) models.DeductBalanceResponse
}

func NewWalletService(consumer rabbitmq.ConsumerInterface, publisher rabbitmq.PublisherInterface) WalletServiceInterface {
	return &WalletService{
		consumer:  consumer,
		publisher: publisher,
	}
}

func (s *WalletService) GetUserBalance(request *models.GetUserBalanceRequest, correlationId string) models.GetUserBalanceResponse {
	// push to request user
	getUserBalanceReplyKey := "wallet.balance.get.request.reply"
	getUserRequestBody, _ := json.Marshal(&request)
	s.publisher.Push("wallet.balance.get.request", getUserBalanceReplyKey, getUserRequestBody, correlationId)

	// waiting for reply
	res, err := s.consumer.WaitReply(getUserBalanceReplyKey, "disburse-consumer", correlationId)

	if err != nil {
		// throw internal server error
		return models.GetUserBalanceResponse{
			Status:    constants.Error,
			ErrorCode: constants.InternalServerError,
		}
	}

	var walletRes models.GetUserBalanceResponse
	json.Unmarshal(res.Body, &walletRes)

	return walletRes
}

func (s *WalletService) DeductUserBalance(request *models.DeductBalanceRequest, correlationId string) models.DeductBalanceResponse {
	// push to request deduct balance
	deductBalanceReplyKey := "wallet.balance.deduct.requestreply"
	deductBalanceRequestBody, _ := json.Marshal(&request)
	s.publisher.Push("wallet.balance.deduct.request", deductBalanceReplyKey, deductBalanceRequestBody, correlationId)

	// waiting for reply
	res, err := s.consumer.WaitReply(deductBalanceReplyKey, "disburse-consumer", correlationId)

	if err != nil {
		// throw internal server error
		return models.DeductBalanceResponse{
			Status:    constants.Error,
			ErrorCode: constants.InternalServerError,
		}
	}

	var balanceRes models.DeductBalanceResponse
	json.Unmarshal(res.Body, &balanceRes)

	return balanceRes
}

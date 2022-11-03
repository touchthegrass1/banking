package services

import (
	"github.com/dopefresh/banking/golang/banking/src/models"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"go.uber.org/zap"
)

type TransactionService struct {
	repository repositories.TransactionRepository
	log        *zap.Logger
}

func (service TransactionService) GetClientTransactions(inn string) ([]models.Transaction, error) {
	return service.repository.GetClientTransactions(inn)
}

func (service TransactionService) GetTransactionById(transactionId int64) (models.Transaction, error) {
	return service.repository.GetTransactionById(transactionId)
}

package services

import (
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"go.uber.org/zap"
)

type TransactionService struct {
	repository repositories.TransactionRepository
	log        *zap.Logger
}

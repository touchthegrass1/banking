package services

import (
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"go.uber.org/zap"
)

func ProvideClientService(repository repositories.ClientRepository, log *zap.Logger) ClientService {
	return ClientService{repository, log}
}

func ProvideTransactionService(repository repositories.TransactionRepository, log *zap.Logger) TransactionService {
	return TransactionService{repository, log}
}

func ProvideCardService(repository repositories.CardRepository, log *zap.Logger) CardService {
	return CardService{repository, log}
}

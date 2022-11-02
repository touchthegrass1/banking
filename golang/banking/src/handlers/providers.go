package handlers

import (
	"github.com/dopefresh/banking/golang/banking/src/services"
	"go.uber.org/zap"
)

func ProvideClientHandler(clientService services.ClientService, log *zap.Logger) ClientHandler {
	return ClientHandler{clientService, log}
}

func ProvideTransactionHandler(transactionService services.TransactionService, log *zap.Logger) TransactionHandler {
	return TransactionHandler{transactionService, log}
}

func ProvideCardHandler(service services.CardService, log *zap.Logger) CardHandler {
	return CardHandler{service, log}
}

package handlers

import (
	"github.com/dopefresh/banking/golang/banking/src/services"
	"go.uber.org/zap"
)

func ProvideClientHandler(clientService services.ClientService, cardPermissionService services.CardPermissionService, log *zap.Logger) ClientHandler {
	return ClientHandler{clientService, cardPermissionService, log}
}

func ProvideTransactionHandler(transactionService services.TransactionService, log *zap.Logger) TransactionHandler {
	return TransactionHandler{transactionService, log}
}

func ProvideCardHandler(service services.CardService, cardPermissionService services.CardPermissionService, log *zap.Logger) CardHandler {
	return CardHandler{service, cardPermissionService, log}
}

package handlers

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dopefresh/banking/golang/banking/src/services"
	"go.uber.org/zap"
)

func ProvideClientHandler(kafkaService services.KafkaService, clientService services.ClientService, cardPermissionService services.CardPermissionService, transactionService services.TransactionService, log *zap.Logger) ClientHandler {
	return ClientHandler{kafkaService, clientService, cardPermissionService, transactionService, log}
}

func ProvideTransactionHandler(transactionService services.TransactionService, log *zap.Logger) TransactionHandler {
	return TransactionHandler{transactionService, log}
}

func ProvideCardHandler(service services.CardService, cardPermissionService services.CardPermissionService, log *zap.Logger) CardHandler {
	return CardHandler{service, cardPermissionService, log}
}

func ProvideKafkaHandler(consumer *kafka.Consumer, log *zap.Logger) KafkaHandler {
	return KafkaHandler{consumer, log}
}

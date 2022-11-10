package services

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dopefresh/banking/golang/banking/src/models"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"go.uber.org/zap"
)

type KafkaTransactionService struct {
	repository *repositories.KafkaTransactionRepository
	log        *zap.Logger
}

type KafkaService interface {
	SendMessage(models.Transaction)
}

func (service KafkaTransactionService) SendMessage(transaction models.Transaction) {
	err := service.repository.SendMessage(transaction)
	if err != nil {
		service.log.Error("Error sending kafka message", zap.Error(err))
	}
}

func HandleKafkaEvents(producer *kafka.Producer, logger *zap.Logger) {
	for e := range producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				logger.Error("Failed to deliver message: ", zap.Error(ev.TopicPartition.Error))
			} else {
				logger.Info("Produced event",
					zap.String("topic", *ev.TopicPartition.Topic),
					zap.String("key", string(ev.Key)),
					zap.String("value", string(ev.Value)),
				)
			}
		}
	}
}

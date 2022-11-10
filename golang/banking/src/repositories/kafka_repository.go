package repositories

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dopefresh/banking/golang/banking/src/models"
)

type KafkaTransactionRepository struct {
	producer     *kafka.Producer
	deliveryChan chan kafka.Event
}

func (repository *KafkaTransactionRepository) SendMessage(transaction models.Transaction) error {
	topic := "transactions"
	value, err := json.Marshal(transaction)
	if err != nil {
		return err
	}
	err = repository.producer.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          value,
		},
		repository.deliveryChan,
	)
	return err
}

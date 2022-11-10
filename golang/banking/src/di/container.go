package di

import (
	"bufio"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dopefresh/banking/golang/banking/src/database_layer"
	"github.com/dopefresh/banking/golang/banking/src/handlers"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"github.com/dopefresh/banking/golang/banking/src/services"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	producer *kafka.Producer
	consumer *kafka.Consumer
	db       *gorm.DB
	Log      *zap.Logger
}

func (container Container) GetClientHandler(kafkaService services.KafkaService) handlers.ClientHandler {
	service := container.GetClientService()
	logger := utils.ProvideLogger()
	cardPermissionService := container.GetCardPermissionService()
	transactionService := container.GetTransactionService()
	return handlers.ProvideClientHandler(kafkaService, service, cardPermissionService, transactionService, logger)
}

func (container Container) GetCardHandler() handlers.CardHandler {
	cardService := container.GetCardService()
	cardPermissionService := container.GetCardPermissionService()
	logger := utils.ProvideLogger()
	return handlers.ProvideCardHandler(cardService, cardPermissionService, logger)
}

func (container Container) GetTransactionHandler() handlers.TransactionHandler {
	service := container.GetTransactionService()
	logger := utils.ProvideLogger()
	return handlers.ProvideTransactionHandler(service, logger)
}

func (container Container) GetKafkaHandler() handlers.KafkaHandler {
	consumer := container.GetKafkaConsumer()
	logger := utils.ProvideLogger()
	return handlers.ProvideKafkaHandler(consumer, logger)
}

func (container Container) GetKafkaTransactionService() services.KafkaTransactionService {
	repository := container.GetKafkaTransactionRepository()
	logger := utils.ProvideLogger()
	return *services.ProvideKafkaTransactionService(&repository, logger)
}

func (container Container) GetJWTService() services.JWTService {
	pubKeyString, exists := os.LookupEnv("PUBLIC_KEY")
	if !exists {
		panic("Didn't provide public key for checking jwt. Env var PUBLIC_KEY")
	}
	logger := utils.ProvideLogger()
	return services.ProvideJWTService(string(pubKeyString), logger)
}

func (container Container) GetCardPermissionService() services.CardPermissionService {
	clientRepository := container.GetClientRepository()
	cardRepository := container.GetCardRepository()
	logger := utils.ProvideLogger()
	return services.ProvideCardPermissionService(clientRepository, cardRepository, logger)
}

func (container Container) GetClientService() services.ClientService {
	repository := container.GetClientRepository()
	logger := utils.ProvideLogger()
	return services.ProvideClientService(repository, logger)
}

func (container Container) GetCardService() services.CardService {
	repository := container.GetCardRepository()
	logger := utils.ProvideLogger()
	return services.ProvideCardService(repository, logger)
}

func (container Container) GetTransactionService() services.TransactionService {
	repository := container.GetTransactionRepository()
	logger := utils.ProvideLogger()
	return services.ProvideTransactionService(repository, logger)
}

func (container Container) GetKafkaTransactionRepository() repositories.KafkaTransactionRepository {
	producer := container.GetKafkaProducer()
	deliveryChan := make(chan kafka.Event, 10000)
	return *repositories.ProvideKafkaTransactionRepository(producer, deliveryChan)
}

func (container Container) GetClientRepository() repositories.ClientRepository {
	db := container.GetDB()
	logger := utils.ProvideLogger()
	return *repositories.ProvideClientRepository(db, logger)
}

func (container Container) GetCardRepository() repositories.CardRepository {
	db := container.GetDB()
	logger := utils.ProvideLogger()
	return *repositories.ProvideCardRepository(db, logger)
}

func (container Container) GetTransactionRepository() repositories.TransactionRepository {
	db := container.GetDB()
	logger := utils.ProvideLogger()
	return *repositories.ProvideTransactionRepository(db, logger)
}

func (container Container) GetDB() *gorm.DB {
	if container.db != nil {
		return container.db
	}
	logger := utils.ProvideLogger()
	dbparams := database_layer.ProvideDBParams()
	gormConfig := database_layer.ProvideGormConfig()
	database, err := database_layer.ProvideDB(dbparams, logger, gormConfig)
	if err != nil {
		panic(err)
	}
	container.db = database
	return database
}

func (container Container) GetKafkaProducer() *kafka.Producer {
	if container.producer != nil {
		return container.producer
	}
	config := container.GetKafkaConfig()
	producer, err := kafka.NewProducer(&config)
	if err != nil {
		panic(err)
	}
	go services.HandleKafkaEvents(producer, container.Log)
	return producer
}

func (container Container) GetKafkaConsumer() *kafka.Consumer {
	if container.consumer != nil {
		return container.consumer
	}
	config := container.GetKafkaConfig()
	config["auto.offset.reset"] = "earliest"
	config["group.id"] = "test-consumer-group"

	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		panic(err)
	}
	topic := "transactions"
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(err)
	}
	return consumer
}

func (container Container) GetKafkaConfig() kafka.ConfigMap {
	m := make(map[string]kafka.ConfigValue)

	file, err := os.Open("kafka.properties")
	if err != nil {
		container.Log.Error("Error getting kafka config", zap.Error(err))
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}
		kv := strings.Split(line, "=")
		key := strings.TrimSpace(kv[0])
		value := strings.TrimSpace(kv[1])
		m[key] = value
	}
	return m
}

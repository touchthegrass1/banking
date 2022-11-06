package di

import (
	"os"

	"github.com/dopefresh/banking/golang/banking/src/database_layer"
	"github.com/dopefresh/banking/golang/banking/src/handlers"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"github.com/dopefresh/banking/golang/banking/src/services"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"gorm.io/gorm"
)

type Container struct {
	db *gorm.DB
}

func (container Container) GetClientHandler() handlers.ClientHandler {
	service := container.GetClientService()
	logger := utils.ProvideLogger()
	cardPermissionService := container.GetCardPermissionService()
	return handlers.ProvideClientHandler(service, cardPermissionService, logger)
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

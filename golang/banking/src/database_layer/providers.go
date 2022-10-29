package database_layer

import (
	"os"
	"strings"

	"github.com/dopefresh/banking/golang/banking/src/utils"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func InitializeDB() *gorm.DB {
	logger := utils.ProvideLogger()
	dbparams := ProvideDBParams()
	database, err := ProvideDB(dbparams, logger)
	if err != nil {
		panic(err)
	}
	return database
}

func ProvideDB(dbparams DatabaseParams, logger *zap.Logger) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbparams.MasterConnectionURI), &gorm.Config{})
	replicas := make([]gorm.Dialector, len(dbparams.ReplicasConnectionURIs))
	for index, value := range dbparams.ReplicasConnectionURIs {
		replicas[index] = postgres.Open(value)
	}

	if err != nil {
		logger.Error("Failed to connect to postgresql using", zap.String("connection uri", dbparams.MasterConnectionURI))
		return nil, err
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: replicas,
		Policy:   dbresolver.RandomPolicy{},
	}))
	return db, nil
}

func ProvideDBParams() DatabaseParams {
	masterConnectionURI, presence := os.LookupEnv("POSTGRESQL_MASTER_CONNECTION_URI")
	if !presence {
		panic("Didn't provide postgresql master connection uri")
	}
	replicaConnectionURIs, presence := os.LookupEnv("POSTGRESQL_REPLICAS_CONNECTION_URIs")
	if !presence {
		return DatabaseParams{MasterConnectionURI: masterConnectionURI, ReplicasConnectionURIs: []string{}}
	}
	replicas := strings.Split(replicaConnectionURIs, ",")
	return DatabaseParams{MasterConnectionURI: masterConnectionURI, ReplicasConnectionURIs: replicas}
}

type DatabaseParams struct {
	MasterConnectionURI    string
	ReplicasConnectionURIs []string
}

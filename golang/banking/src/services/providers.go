package services

import (
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"github.com/lestrrat-go/jwx/v2/jwk"
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

func ProvideJWTService(publicKey string, log *zap.Logger) JWTService {
	key, err := jwk.ParseKey([]byte(publicKey), jwk.WithPEM(true))
	if err != nil {
		panic(err)
	}

	service := JWTService{key, log}
	return service
}

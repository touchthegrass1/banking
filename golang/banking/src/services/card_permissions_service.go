package services

import (
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"go.uber.org/zap"
)

type CardPermissionService struct {
	clientRepository repositories.ClientRepository
	cardRepository   repositories.CardRepository
	log              *zap.Logger
}

func (service CardPermissionService) CheckCanUseCard(cardNumber string, userId int64) bool {
	client, err := service.clientRepository.GetClientByUserIdWithCards(userId)
	if err != nil {
		return false
	}
	for i := 0; i < len(client.Cards); i++ {
		if client.Cards[i].CardId == cardNumber {
			return true
		}
	}
	return false
}

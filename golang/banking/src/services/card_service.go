package services

import (
	"github.com/dopefresh/banking/golang/banking/src/models"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"go.uber.org/zap"
)

type CardService struct {
	Repository repositories.CardRepository
	Log        *zap.Logger
}

func (service CardService) CreateCard(userId int64, card models.Card) error {
	return service.Repository.AddCard(userId, card)
}

func (service CardService) GetCardByNumber(cardNumber string) (models.Card, error) {
	return service.Repository.GetCard(cardNumber)
}

func (service CardService) UpdateCard(cardNumber string, card models.CardUpdate) error {
	return service.Repository.UpdateCard(cardNumber, card)
}

func (service CardService) DeleteCard(cardNumber string) error {
	return service.Repository.DeleteCard(cardNumber)
}

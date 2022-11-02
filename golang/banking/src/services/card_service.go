package services

import (
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"go.uber.org/zap"
)

type CardService struct {
	Repository repositories.CardRepository
	Log        *zap.Logger
}

func (service CardService) CreateCard() {

}

func (service CardService) GetCard() {

}

func (service CardService) UpdateCard() {

}

func (service CardService) DeleteCard() {

}

package repositories

import (
	"github.com/dopefresh/banking/golang/banking/src/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CardRepository struct {
	Db  *gorm.DB
	Log *zap.Logger
}

func (repository CardRepository) GetDB() *gorm.DB {
	return repository.Db
}

func (repository CardRepository) AddCard(card models.Card) {
	repository.GetDB().Create(&card)
}

func (repository CardRepository) GetCard() {

}

func (repository CardRepository) UpdateCard() {

}

func (repository CardRepository) DeleteCard() {

}

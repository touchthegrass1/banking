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

func (repository CardRepository) AddCard(card models.Card) error {
	return repository.GetDB().Create(&card).Error
}

func (repository CardRepository) GetCard(cardNumber string) (models.Card, error) {
	var card models.Card
	err := repository.GetDB().Where("card_id = ?", cardNumber).Find(&card).Error
	return card, err
}

func (repository CardRepository) UpdateCard(cardNumber string, card models.CardUpdate) error {
	err := repository.GetDB().Where("card_id = ?", cardNumber).Updates(card).Error
	return err
}

func (repository CardRepository) DeleteCard(cardNumber string) error {
	err := repository.GetDB().Delete(&models.Card{}, cardNumber).Error
	return err
}

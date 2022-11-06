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

func (repository CardRepository) AddCard(userId int64, card models.Card) error {
	db := repository.GetDB()

	var client models.Client
	err := db.Model(&models.Client{}).Where("user_id = ?", userId).Select("client_id").First(&client).Error
	if err != nil {
		return err
	}
	card.ClientId = client.ClientId
	return repository.GetDB().Create(&card).Error
}

func (repository CardRepository) GetCard(cardNumber string) (models.Card, error) {
	var card models.Card
	err := repository.GetDB().Where("card_id = ?", cardNumber).First(&card).Error
	return card, err
}

func (repository CardRepository) UpdateCard(cardNumber string, card models.CardUpdate) error {
	err := repository.GetDB().Where("card_id = ?", cardNumber).Updates(card).Error
	return err
}

func (repository CardRepository) DeleteCard(cardNumber string) error {
	err := repository.GetDB().Where("card_id = ?", cardNumber).Delete(&models.Card{}).Error
	return err
}

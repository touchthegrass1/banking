package repositories

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CardRepository struct {
	Db  *gorm.DB
	Log *zap.Logger
}

func (repository CardRepository) AddCard() {
}

func (repository CardRepository) GetCard() {

}

func (repository CardRepository) UpdateCard() {

}

func (repository CardRepository) DeleteCard() {

}

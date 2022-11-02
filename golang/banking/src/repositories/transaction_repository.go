package repositories

import (
	"database/sql"

	"github.com/dopefresh/banking/golang/banking/src/database_layer"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	Db  *gorm.DB
	Log *zap.Logger
}

func (transactionRepository TransactionRepository) GetDB() *gorm.DB {
	return transactionRepository.Db
}

func (transactionRepository TransactionRepository) GetClientTransactions(inn string) ([]database_layer.Transaction, error) {
	db := transactionRepository.GetDB()
	var cards []string
	err := db.Model(&database_layer.Card{}).Where(
		"client_id = (?)",
		db.Table("client").Where("inn = @inn", sql.Named("inn", inn)).Select("client_id"),
	).Select("card_id").Find(&cards).Error

	if err != nil {
		return nil, err
	}

	var transactions []database_layer.Transaction
	err = db.Where(
		"card_from_id IN @cardIds OR card_to_id IN @cardIds OR card_id IN @cardIds",
		sql.Named("cardIds", cards),
	).Find(&transactions).Error

	return transactions, err
}

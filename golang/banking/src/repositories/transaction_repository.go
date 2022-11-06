package repositories

import (
	"database/sql"
	"time"

	"github.com/dopefresh/banking/golang/banking/src/models"
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

func (transactionRepository TransactionRepository) GetClientTransactions(userId int64) ([]models.Transaction, error) {
	db := transactionRepository.GetDB()
	var cards []string
	err := db.Model(&models.Card{}).Where(
		"client_id = (?)",
		db.Table("client").Where("userId = ?", userId).Select("client_id"),
	).Select("card_id").Find(&cards).Error

	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	err = db.Where(
		"card_from_id IN @cardIds OR card_to_id IN @cardIds OR card_id IN @cardIds",
		sql.Named("cardIds", cards),
	).Find(&transactions).Error

	return transactions, err
}

func (repository TransactionRepository) GetTransactionById(transactionId int64) (models.Transaction, error) {
	var transaction models.Transaction
	err := repository.GetDB().Model(&models.Transaction{}).Where("transaction_id = ?", transactionId).Find(&transaction).Error
	return transaction, err
}

// Unused function. Django admin has this functionality
func (repository TransactionRepository) GetTransactionByDatetimeRange(datetimeBegin time.Time, datetimeEnd time.Time) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := repository.GetDB().Where(
		"transaction_datetime BETWEEN @datetimeBegin AND @datetimeEnd",
		sql.Named("datetimeBegin", datetimeBegin),
		sql.Named("datetimeEnd", datetimeEnd),
	).Find(&transactions).Error
	return transactions, err
}

package repositories

import (
	"fmt"

	"github.com/dopefresh/banking/golang/banking/src/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ClientRepository struct {
	Db  *gorm.DB
	Log *zap.Logger
}

func (clientRepository ClientRepository) GetDB() *gorm.DB {
	return clientRepository.Db
}

func (clientRepository ClientRepository) GetClientByUserId(userId int64) (models.Client, error) {
	var client models.Client
	err := clientRepository.GetDB().Table("client").Where("user_id = ?", userId).Find(&client).Error
	return client, err
}

func (clientRepository ClientRepository) GetClientByUserIdWithCards(userId int64) (models.Client, error) {
	var client models.Client
	err := clientRepository.GetDB().Model(&models.Client{}).Preload("Cards").Where("user_id = ?", userId).Find(&client).Error
	return client, err
}

func (clientRepository ClientRepository) UpdateClientByUserId(userId int64, client models.ClientUpdate) error {
	return clientRepository.GetDB().Model(&models.Client{}).Where("user_id = ?", userId).Updates(client).Error
}

func (repository ClientRepository) TransferMoney(transfer models.Transfer) (int64, error) {
	var transactionId int64 = 0

	err := repository.GetDB().Transaction(func(tx *gorm.DB) error {
		tx.Table("card").Where("card_id = ?", transfer.CardFromId).Update("balance", gorm.Expr("balance - ?", transfer.Summ))
		tx.Table("card").Where("card_id = ?", transfer.CardToId).Update("balance", gorm.Expr("balance + ?", transfer.Summ))
		transaction := models.Transaction{
			TransactionType: "transfer",
			CardFromId:      transfer.CardFromId,
			CardToId:        transfer.CardToId,
			Summ:            transfer.Summ,
		}
		result := tx.Create(&transaction)
		if result.Error != nil {
			repository.Log.Error("Error transferring money", zap.Error(result.Error))
			return fmt.Errorf("error transferring money %v: %w", transaction, result.Error)
		}
		transactionId = transaction.TransactionId
		return nil
	})
	return transactionId, err
}

func (repository ClientRepository) DepositMoney(deposit models.Deposit) (int64, error) {
	var transactionId int64 = 0

	err := repository.GetDB().Transaction(func(tx *gorm.DB) error {
		tx.Table("card").Where("card_id = ?", deposit.CardId).Update("balance", gorm.Expr("balance + ?", deposit.Summ))
		transaction := models.Transaction{
			TransactionType: "deposit",
			CardId:          deposit.CardId,
			Summ:            deposit.Summ,
		}
		result := tx.Create(&transaction)
		if result.Error != nil {
			repository.Log.Error("Error depositing money", zap.Error(result.Error))
			return fmt.Errorf("error depositing money %v: %w", transaction, result.Error)
		}
		transactionId = transaction.TransactionId
		return nil
	})
	return transactionId, err
}

func (repository ClientRepository) WithdrawMoney(withdraw models.Withdraw) (int64, error) {
	var transactionId int64 = 0

	err := repository.GetDB().Transaction(func(tx *gorm.DB) error {
		tx.Table("card").Where("card_id = ?", withdraw.CardId).Update("balance", gorm.Expr("balance - ?", withdraw.Summ))
		transaction := models.Transaction{
			TransactionType: "withdraw",
			CardId:          withdraw.CardId,
			Summ:            withdraw.Summ,
		}
		result := tx.Create(&transaction)

		if result.Error != nil {
			repository.Log.Error("Error money withdraw", zap.Error(result.Error))
			return fmt.Errorf("error withdrawing money %v: %w", transaction, result.Error)
		}

		transactionId = transaction.TransactionId
		return nil
	})
	return transactionId, err
}

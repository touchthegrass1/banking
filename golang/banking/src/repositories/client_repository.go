package repositories

import (
	"database/sql"

	"github.com/dopefresh/banking/golang/banking/src/database_layer"
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

func (clientRepository ClientRepository) GetClientByInn(inn string) (database_layer.Client, error) {
	var client database_layer.Client
	err := clientRepository.GetDB().Table("client").Where("inn = @inn", sql.Named("inn", inn)).Find(&client).Error
	return client, err
}

func (clientRepository ClientRepository) GetClientByInnWithCards(inn string) (database_layer.Client, error) {
	var client database_layer.Client
	err := clientRepository.GetDB().Model(&database_layer.Client{}).Preload("Cards").Where("inn = @inn", sql.Named("inn", inn)).Find(&client).Error
	return client, err
}

func (clientRepository ClientRepository) UpdateClientByInn(inn string, client models.ClientUpdate) error {
	clientRepository.GetDB().Model(&database_layer.Client{}).Where("inn = ?", inn).Updates(client)
	clientRepository.GetDB().Save(&client)
	return nil
}

func (clientRepository ClientRepository) TransferMoney(transfer models.Transfer) error {
	err := clientRepository.GetDB().Transaction(func(tx *gorm.DB) error {
		tx.Table("card").Where("card_id = ?", transfer.CardFromId).Update("balance", gorm.Expr("balance - ?", transfer.Summ))
		tx.Table("card").Where("card_id = ?", transfer.CardToId).Update("balance", gorm.Expr("balance + ?", transfer.Summ))
		transaction := database_layer.Transaction{
			TransactionType: "transfer",
			CardFromId:      transfer.CardFromId,
			CardToId:        transfer.CardToId,
			Summ:            transfer.Summ,
		}
		tx.Create(&transaction)
		return nil
	})
	return err
}

func (clientRepository ClientRepository) DepositMoney(deposit models.Deposit) error {
	err := clientRepository.GetDB().Transaction(func(tx *gorm.DB) error {
		tx.Table("card").Where("card_id = ?", deposit.CardId).Update("balance", gorm.Expr("balance + ?", deposit.Summ))
		transaction := database_layer.Transaction{
			TransactionType: "deposit",
			CardId:          deposit.CardId,
			Summ:            deposit.Summ,
		}
		tx.Create(&transaction)
		return nil
	})
	return err
}

func (clientRepository ClientRepository) WithdrawMoney(withdraw models.Withdraw) error {
	err := clientRepository.GetDB().Transaction(func(tx *gorm.DB) error {
		tx.Table("card").Where("card_id = ?", withdraw.CardId).Update("balance", gorm.Expr("balance - ?", withdraw.Summ))
		transaction := database_layer.Transaction{
			TransactionType: "withdraw",
			CardId:          withdraw.CardId,
			Summ:            withdraw.Summ,
		}
		tx.Create(&transaction)
		return nil
	})
	return err
}

package repositories

import (
	"database/sql"

	"github.com/dopefresh/banking/golang/banking/src/database_layer"
	"github.com/shopspring/decimal"
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
	clientRepository.GetDB().Table("client").Where("inn = @inn", sql.Named("inn", inn)).Find(&client)
	return client, nil
}

func (clientRepository ClientRepository) UpdateClient(client database_layer.Client) error {
	clientRepository.GetDB().Save(&client)
	return nil
}

func (clientRepository ClientRepository) TransferMoney(cardIdFrom string, cardIdTo string, sum decimal.Decimal) {
	clientRepository.GetDB().Transaction(func(tx *gorm.DB) error {
		tx.Exec(
			"UPDATE card SET balance = balance - @sum WHERE card_id = @cardIdFrom",
			sql.Named("sum", sum),
			sql.Named("cardIdFrom", cardIdFrom),
		)
		tx.Exec(
			"UPDATE card SET balance = balance + @sum WHERE card_id = @cardIdTo",
			sql.Named("sum", sum),
			sql.Named("cardIdTo", cardIdTo),
		)
		// TODO: Add insert to transaction
		return nil
	})
}

func ProvideClientRepository(db *gorm.DB, log *zap.Logger) *ClientRepository {
	clientRepository := ClientRepository{Db: db, Log: log}
	return &clientRepository
}

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

func (clientRepository ClientRepository) UpdateClient(oldInn string, client database_layer.Client) error {
	clientRepository.GetDB().Exec(
		"UPDATE client SET registration_address = @registration_address, residential_address = @residential_address, client_type = @client_type, ogrn = @ogrn, inn = @inn, kpp = @kpp WHERE inn = @old_inn",
		sql.Named("registration_address", client.RegistrationAddress),
		sql.Named("residential_address", client.ResidentialAddress),
		sql.Named("client_type", client.ClientType),
		sql.Named("ogrn", client.Ogrn),
		sql.Named("inn", client.Inn),
		sql.Named("kpp", client.Kpp),
		sql.Named("old_inn", oldInn),
	)

	return nil
}

func (clientRepository ClientRepository) AddNewClient(client database_layer.Client) error {
	clientRepository.GetDB().Create(&client)
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

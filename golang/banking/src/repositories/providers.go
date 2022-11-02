package repositories

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ProvideClientRepository(db *gorm.DB, log *zap.Logger) *ClientRepository {
	return &ClientRepository{Db: db, Log: log}
}

func ProvideTransactionRepository(db *gorm.DB, log *zap.Logger) *TransactionRepository {
	return &TransactionRepository{Db: db, Log: log}
}

func ProvideCardRepository(db *gorm.DB, log *zap.Logger) *CardRepository {
	return &CardRepository{Db: db, Log: log}
}

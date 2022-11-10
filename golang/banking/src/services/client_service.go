package services

import (
	"github.com/dopefresh/banking/golang/banking/src/models"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"go.uber.org/zap"
)

type ClientService struct {
	repository repositories.ClientRepository
	log        *zap.Logger
}

func (clientService ClientService) GetClientByUserId(userId int64) (models.Client, error) {
	client, err := clientService.repository.GetClientByUserId(userId)
	return client, err
}

func (clientService ClientService) UpdateClientByUserId(userId int64, client models.ClientUpdate) error {
	err := clientService.repository.UpdateClientByUserId(userId, client)
	return err
}

func (service ClientService) DepositMoney(deposit models.Deposit) (int64, error) {
	transactionId, err := service.repository.DepositMoney(deposit)
	return transactionId, err
}

func (service ClientService) WithdrawMoney(withdraw models.Withdraw) (int64, error) {
	transactionId, err := service.repository.WithdrawMoney(withdraw)
	return transactionId, err
}

func (clientService ClientService) TransferMoney(transfer models.Transfer) (int64, error) {
	transactionId, err := clientService.repository.TransferMoney(transfer)
	return transactionId, err
}

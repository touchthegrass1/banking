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

func (clientService ClientService) GetClientByInn(inn string) (models.Client, error) {
	client, err := clientService.repository.GetClientByInn(inn)
	return client, err
}

func (clientService ClientService) UpdateClientByInn(inn string, client models.ClientUpdate) error {
	err := clientService.repository.UpdateClientByInn(inn, client)
	return err
}

func (service ClientService) DepositMoney(deposit models.Deposit) error {
	err := service.repository.DepositMoney(deposit)
	return err
}

func (service ClientService) WithdrawMoney(withdraw models.Withdraw) error {
	err := service.repository.WithdrawMoney(withdraw)
	return err
}

func (clientService ClientService) TransferMoney(transfer models.Transfer) error {
	err := clientService.repository.TransferMoney(transfer)
	return err
}

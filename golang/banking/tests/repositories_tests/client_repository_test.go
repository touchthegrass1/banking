package repositories_tests

import (
	"database/sql"
	"testing"

	"github.com/dopefresh/banking/golang/banking/src/di"
	"github.com/dopefresh/banking/golang/banking/src/models"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientRepositoryTestSuite struct {
	suite.Suite
	clientRepository repositories.ClientRepository
	clients          []models.Client
	users            []models.User
}

func TestClientRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ClientRepositoryTestSuite))
}

func (suite *ClientRepositoryTestSuite) SetupSuite() {
	container := di.Container{}
	db := container.GetDB()
	log := utils.ProvideLogger()
	suite.clientRepository = *repositories.ProvideClientRepository(db, log)
	suite.setupData()
}

func (suite *ClientRepositoryTestSuite) TearDownSuite() {
	suite.clientRepository.Db.Exec(`SELECT truncate_tables('banking_db_user', '{"client","card","contract","credit","payment_schedule","transaction","user"}');`)
}

func (suite *ClientRepositoryTestSuite) setupData() {
	suite.clientRepository.Db.Exec("CALL p_setup_test_data()")
}

func (suite *ClientRepositoryTestSuite) TestClientRepositoryGetClientByInn() {
	var client models.Client
	suite.clientRepository.GetDB().Take(&client)

	foundClient, err := suite.clientRepository.GetClientByUserId(client.UserId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), client, foundClient)
}

func (suite *ClientRepositoryTestSuite) TestClientRepositoryUpdateClient() {
	var client models.Client
	suite.clientRepository.GetDB().Take(&client)
	clientUpdate := models.ClientUpdate{
		FirstName:           "Vasilii",
		LastName:            "Popov",
		Phone:               "89996458671",
		RegistrationAddress: "Moscow, Lva Tolstogo 1",
		ResidentialAddress:  "Moscow, Lva Tolstogo 1",
		ClientType:          "jp",
	}
	err := suite.clientRepository.UpdateClientByUserId(client.UserId, clientUpdate)
	assert.Nil(suite.T(), err)

	clientFound, err := suite.clientRepository.GetClientByUserId(client.UserId)
	assert.Equal(suite.T(), clientFound.RegistrationAddress, "Moscow, Lva Tolstogo 1")
	assert.Equal(suite.T(), clientFound.ResidentialAddress, "Moscow, Lva Tolstogo 1")
	assert.Equal(suite.T(), clientFound.ClientType, models.ClientType("jp"))
}

func (suite *ClientRepositoryTestSuite) TestClientRepositoryTransfer() {
	var client1, client2 models.Client
	err := suite.clientRepository.GetDB().Model(&models.Client{}).Preload("Cards").Where("client_id = @clientId", sql.Named("clientId", 1)).Find(&client1).Error
	assert.Nil(suite.T(), err)
	err = suite.clientRepository.GetDB().Model(&models.Client{}).Preload("Cards").Where("client_id = @clientId", sql.Named("clientId", 2)).Find(&client2).Error
	assert.Nil(suite.T(), err)

	want := client1.Cards[0].Balance.Add(client2.Cards[0].Balance)

	transfer1 := models.Transfer{
		CardFromId: client1.Cards[0].CardId,
		CardToId:   client2.Cards[0].CardId,
		Summ:       client1.Cards[0].Balance,
	}
	err = suite.clientRepository.TransferMoney(transfer1)
	assert.Nil(suite.T(), err)

	foundClient1, err := suite.clientRepository.GetClientByUserIdWithCards(client1.UserId)
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), 0, foundClient1.Cards[0].Balance.Cmp(decimal.NewFromFloat(0)))
	foundClient2, err := suite.clientRepository.GetClientByUserIdWithCards(client2.UserId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, foundClient2.Cards[0].Balance.Cmp(want))

	transfer2 := models.Transfer{
		CardFromId: foundClient1.Cards[0].CardId,
		CardToId:   foundClient2.Cards[0].CardId,
		Summ:       decimal.NewFromInt(1),
	}

	err = suite.clientRepository.TransferMoney(transfer2)
	assert.NotNil(suite.T(), err)
}

func (suite *ClientRepositoryTestSuite) TestClientRepositoryDeposit() {
	var client models.Client
	err := suite.clientRepository.GetDB().Model(&models.Client{}).Preload("Cards").Where("client_id = @clientId", sql.Named("clientId", 1)).Find(&client).Error
	assert.Nil(suite.T(), err)

	deposit := models.Deposit{
		CardId: client.Cards[0].CardId,
		Summ:   decimal.NewFromInt(100),
	}
	want := client.Cards[0].Balance.Add(decimal.NewFromInt(100))
	err = suite.clientRepository.DepositMoney(deposit)
	assert.Nil(suite.T(), err)

	foundClient, err := suite.clientRepository.GetClientByUserIdWithCards(client.UserId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), want, foundClient.Cards[0].Balance)
}

func (suite *ClientRepositoryTestSuite) TestClientRepositoryWithdraw() {
	var client models.Client
	err := suite.clientRepository.GetDB().Model(&models.Client{}).Preload("Cards").Where("client_id = @clientId", sql.Named("clientId", 1)).Find(&client).Error
	assert.Nil(suite.T(), err)

	withdraw := models.Withdraw{
		CardId: client.Cards[0].CardId,
		Summ:   client.Cards[0].Balance,
	}
	err = suite.clientRepository.WithdrawMoney(withdraw)
	assert.Nil(suite.T(), err)

	foundClient, err := suite.clientRepository.GetClientByUserIdWithCards(client.UserId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, foundClient.Cards[0].Balance.Cmp(decimal.NewFromFloat(0)))

	withdraw2 := models.Withdraw{
		CardId: client.Cards[0].CardId,
		Summ:   decimal.NewFromInt(1),
	}
	err = suite.clientRepository.WithdrawMoney(withdraw2)
	assert.NotNil(suite.T(), err)
}

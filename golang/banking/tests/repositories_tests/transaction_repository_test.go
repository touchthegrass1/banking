package repositories_tests

import (
	"database/sql"
	"testing"

	"github.com/dopefresh/banking/golang/banking/src/di"
	"github.com/dopefresh/banking/golang/banking/src/models"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionRepositoryTestSuite struct {
	suite.Suite
	clientRepository      repositories.ClientRepository
	transactionRepository repositories.TransactionRepository
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ClientRepositoryTestSuite))
}

func (suite *TransactionRepositoryTestSuite) SetupSuite() {
	logger := utils.ProvideLogger()
	container := di.Container{Log: logger}
	db := container.GetDB()
	log := utils.ProvideLogger()
	suite.transactionRepository = *repositories.ProvideTransactionRepository(db, log)
	suite.clientRepository = *repositories.ProvideClientRepository(db, log)
	suite.setupData()
}

func (suite *TransactionRepositoryTestSuite) TearDownSuite() {
	suite.transactionRepository.GetDB().Exec(`SELECT truncate_tables('banking_db_user', '{"client","card","contract","credit","payment_schedule","transaction","user"}');`)
}

func (suite *TransactionRepositoryTestSuite) setupData() {
	suite.transactionRepository.GetDB().Exec("CALL p_setup_test_data()")
}

func (suite *TransactionRepositoryTestSuite) TestTransactionRepositoryGetClientTransactions() {
	var client models.Client
	err := suite.clientRepository.GetDB().Model(&models.Client{}).Where("client_id = @clientId", sql.Named("clientId", 1)).Find(&client).Error
	assert.Nil(suite.T(), err)
	transactions, err := suite.transactionRepository.GetClientTransactions(client.UserId)
	assert.Nil(suite.T(), err)

	assert.NotEmpty(suite.T(), transactions)
}

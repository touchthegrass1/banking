package repositories_tests

import (
	"database/sql"
	"testing"

	"github.com/dopefresh/banking/golang/banking/src/database_layer"
	"github.com/dopefresh/banking/golang/banking/src/di"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TransactionRepositoryTestSuite struct {
	suite.Suite
	clientRepository      repositories.ClientRepository
	transactionRepository repositories.TransactionRepository
	clients               []database_layer.Client
	users                 []database_layer.User
}

func TestTransactionRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ClientRepositoryTestSuite))
}

func (suite *TransactionRepositoryTestSuite) SetupSuite() {
	container := di.Container{}
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
	var client database_layer.Client
	err := suite.clientRepository.GetDB().Model(&database_layer.Client{}).Where("client_id = @clientId", sql.Named("clientId", 1)).Find(&client).Error
	assert.Nil(suite.T(), err)
	transactions, err := suite.transactionRepository.GetClientTransactions(client.Inn)
	assert.Nil(suite.T(), err)

	assert.NotEmpty(suite.T(), transactions)
}

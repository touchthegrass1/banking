package handlers_tests

import (
	"fmt"
	"testing"

	"github.com/dopefresh/banking/golang/banking/src/di"
	"github.com/dopefresh/banking/golang/banking/src/handlers"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type TransactionHandlerTestSuite struct {
	suite.Suite
	handler handlers.TransactionHandler
	db      *gorm.DB
	tokens  Tokens
}

func TestTransactionHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionHandlerTestSuite))
}

func (suite *TransactionHandlerTestSuite) SetupSuite() {
	logger := utils.ProvideLogger()
	container := di.Container{Log: logger}
	suite.handler = container.GetTransactionHandler()
	suite.db = container.GetDB()
	suite.setupData()
}

func (suite *TransactionHandlerTestSuite) setupData() {
	tokens, err := RegisterAndGetTokens("4phone1", "4inn1")
	if err != nil {
		panic(fmt.Sprintf("Error occured getting tokens %s", err))
	}
	suite.tokens = tokens
}

func (suite *TransactionHandlerTestSuite) TearDownSuite() {
	suite.db.Exec(`SELECT truncate_tables('banking_db_user', '{"client","card","contract","credit","payment_schedule","transaction","user"}');`)
}

func (suite *TransactionHandlerTestSuite) TestGetTransactions() {
	err := createCard("4card1", decimal.NewFromInt(100), suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	err = deposit("4card1", 100, suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	err = withdraw("4card1", 50, suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	transactions, err := getTransactions(suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, transactions[0].Summ.Cmp(decimal.NewFromInt(100)))
	assert.Equal(suite.T(), 0, transactions[1].Summ.Cmp(decimal.NewFromInt(50)))
	assert.Equal(suite.T(), "4card1", transactions[0].CardId)
	assert.Equal(suite.T(), "4card1", transactions[0].CardId)
	assert.Equal(suite.T(), "", transactions[0].CardFromId)
	assert.Equal(suite.T(), "", transactions[0].CardToId)
	assert.Equal(suite.T(), "", transactions[1].CardFromId)
	assert.Equal(suite.T(), "", transactions[1].CardToId)
}

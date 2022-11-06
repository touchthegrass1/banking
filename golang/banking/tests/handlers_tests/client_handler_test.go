package handlers_tests

import (
	"fmt"
	"testing"

	"github.com/dopefresh/banking/golang/banking/src/di"
	"github.com/dopefresh/banking/golang/banking/src/handlers"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type ClientHandlerTestSuite struct {
	suite.Suite
	handler handlers.ClientHandler
	db      *gorm.DB
	tokens  Tokens
}

func TestClientHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ClientHandlerTestSuite))
}

func (suite *ClientHandlerTestSuite) SetupSuite() {
	container := di.Container{}
	suite.handler = container.GetClientHandler()
	suite.db = container.GetDB()
	suite.setupData()
}

func (suite *ClientHandlerTestSuite) setupData() {
	tokens, err := RegisterAndGetTokens("0phone1", "0inn1")
	if err != nil {
		panic(fmt.Sprintf("Error occured getting tokens %s", err))
	}
	suite.tokens = tokens
}

func (suite *ClientHandlerTestSuite) TearDownSuite() {
	suite.db.Exec(`SELECT truncate_tables('banking_db_user', '{"client","card","contract","credit","payment_schedule","transaction","user"}');`)
}

func (suite *ClientHandlerTestSuite) TestUpdateClient() {
	status, err := updateClient(suite.tokens.AccessToken, "0phone1")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 200, status)

	client, err := getClient(suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Russia, Moscow, Lev Tolstoy street 5", client.RegistrationAddress)
	assert.Equal(suite.T(), "Russia, Moscow, Pushkina street 4", client.ResidentialAddress)
}

func (suite *ClientHandlerTestSuite) TestGetClient() {
	client, err := getClient(suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "0inn1", client.Inn)
}

func (suite *ClientHandlerTestSuite) TestTransfer() {
	err := createCard("0card1", decimal.NewFromInt(100), suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	tokens, err := RegisterAndGetTokens("0phone2", "0inn2")
	assert.Nil(suite.T(), err)

	err = createCard("0card2", decimal.NewFromInt(200), tokens.AccessToken)
	assert.Nil(suite.T(), err)

	// runtime.Breakpoint()
	err = transfer("0card1", "0card2", 100, suite.tokens.AccessToken)

	assert.Nil(suite.T(), err)
}

func (suite *ClientHandlerTestSuite) TestDeposit() {
	err := createCard("0card3", decimal.NewFromInt(1000), suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	// runtime.Breakpoint()
	err = deposit("0card3", 1000, suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)
}

func (suite *ClientHandlerTestSuite) TestWithdraw() {
	// runtime.Breakpoint()
	err := createCard("0card4", decimal.NewFromInt(1000), suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	err = withdraw("0card4", 1000, suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)
}

package handlers_tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/dopefresh/banking/golang/banking/src/di"
	"github.com/dopefresh/banking/golang/banking/src/handlers"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CardHandlerTestSuite struct {
	suite.Suite
	handler handlers.CardHandler
	db      *gorm.DB
	tokens  Tokens
}

func TestCardHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(CardHandlerTestSuite))
}

func (suite *CardHandlerTestSuite) SetupSuite() {
	logger := utils.ProvideLogger()
	container := di.Container{Log: logger}
	suite.handler = container.GetCardHandler()
	suite.db = container.GetDB()
	suite.setupData()
}

func (suite *CardHandlerTestSuite) setupData() {
	tokens, err := RegisterAndGetTokens("1phone1", "1inn1")
	if err != nil {
		panic(fmt.Sprintf("Error occured getting tokens %s", err))
	}
	suite.tokens = tokens
}

func (suite *CardHandlerTestSuite) TearDownSuite() {
	suite.db.Exec(`SELECT truncate_tables('banking_db_user', '{"client","card","contract","credit","payment_schedule","transaction","user"}');`)
}

func (suite *CardHandlerTestSuite) TestCardCreateAndGet() {
	err := createCard("1card1", decimal.NewFromInt(1000), suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	card, err := getCard("1card1", suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, card.Balance.Cmp(decimal.NewFromInt(1000)))
	assert.Equal(suite.T(), "1card1", card.CardId)
}

func (suite *CardHandlerTestSuite) TestCardUpdate() {
	err := createCard("1card2", decimal.NewFromInt(1000), suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	err = updateCard("1card2", "000", suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	card, err := getCard("1card2", suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "000", card.CvcCode)
}

func (suite *CardHandlerTestSuite) TestCardDelete() {
	err := createCard("1card3", decimal.NewFromInt(1000), suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)

	status, err := deleteCard("1card3", suite.tokens.AccessToken)
	assert.Equal(suite.T(), http.StatusOK, status)
	assert.Nil(suite.T(), err)

	statusCode, err := getNoCard("1card3", suite.tokens.AccessToken)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusUnauthorized, statusCode)
}

package repositories_tests

import (
	"testing"
	"time"

	"github.com/dopefresh/banking/golang/banking/src/di"
	"github.com/dopefresh/banking/golang/banking/src/models"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CardRepositoryTestSuite struct {
	suite.Suite
	clientRepository repositories.ClientRepository
	cardRepository   repositories.CardRepository
	clients          []models.Client
	users            []models.User
}

func TestCardRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CardRepositoryTestSuite))
}

func (suite *CardRepositoryTestSuite) SetupSuite() {
	container := di.Container{}
	db := container.GetDB()
	log := utils.ProvideLogger()
	suite.clientRepository = *repositories.ProvideClientRepository(db, log)
	suite.cardRepository = *repositories.ProvideCardRepository(db, log)
	suite.setupData()
}

func (suite *CardRepositoryTestSuite) setupData() {
	suite.clientRepository.Db.Exec("CALL p_setup_test_data()")

	var client1, client2 models.Client
	suite.clientRepository.GetDB().First(&client1, 1)
	suite.clientRepository.GetDB().First(&client2, 2)
	suite.clients = append(suite.clients, client1, client2)

	var user1, user2 models.User
	suite.clientRepository.GetDB().First(&user1, 1)
	suite.clientRepository.GetDB().First(&user2, 2)
	suite.users = append(suite.users, user1, user2)
}

func (suite *CardRepositoryTestSuite) TearDownSuite() {
	suite.clientRepository.Db.Exec(`SELECT truncate_tables('banking_db_user', '{"client","card","contract","credit","payment_schedule","transaction","user"}');`)
}

func (suite *CardRepositoryTestSuite) TestCardRepositoryCardCreateAndGet() {
	assert.Equal(suite.T(), 1, 1)
	card := models.Card{
		CardId:   "100020003000",
		Balance:  decimal.NewFromInt(100),
		ValidTo:  time.Now().Add(time.Hour),
		CvcCode:  "010",
		CardType: "debit",
		Currency: "RUB",
		ClientId: suite.clients[0].ClientId,
	}
	err := suite.cardRepository.AddCard(card)
	assert.Nil(suite.T(), err)

	foundCard, err := suite.cardRepository.GetCard("100020003000")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), card.CardId, foundCard.CardId)
	assert.Equal(suite.T(), 0, card.Balance.Cmp(foundCard.Balance))
	assert.Equal(suite.T(), card.CvcCode, foundCard.CvcCode)
	assert.Equal(suite.T(), card.CardType, foundCard.CardType)
	assert.Equal(suite.T(), card.Currency, foundCard.Currency)
}

func (suite *CardRepositoryTestSuite) TestCardRepositoryUpdate() {
	card := models.CardUpdate{
		ValidTo: time.Now().Add(time.Hour),
		CvcCode: "030",
	}
	err := suite.cardRepository.UpdateCard("100020003000", card)
	assert.Nil(suite.T(), err)

	foundCard, err := suite.cardRepository.GetCard("100020003000")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), card.CvcCode, foundCard.CvcCode)
}

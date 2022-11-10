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
	"gorm.io/gorm"
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
	logger := utils.ProvideLogger()
	container := di.Container{Log: logger}
	db := container.GetDB()
	log := utils.ProvideLogger()
	suite.clientRepository = *repositories.ProvideClientRepository(db, log)
	suite.cardRepository = *repositories.ProvideCardRepository(db, log)
	suite.setupData()
}

func (suite *CardRepositoryTestSuite) setupData() {
	suite.clientRepository.Db.Exec("CALL p_setup_test_data()")

	client1, err := suite.clientRepository.GetClientByUserIdWithCards(1)
	assert.Nil(suite.T(), err)

	client2, err := suite.clientRepository.GetClientByUserIdWithCards(2)
	assert.Nil(suite.T(), err)

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
	card := models.Card{
		CardId:   "2card1",
		Balance:  decimal.NewFromInt(100),
		ValidTo:  time.Now().Add(time.Hour),
		CvcCode:  "010",
		CardType: "debit",
		Currency: "RUB",
		ClientId: suite.clients[0].ClientId,
	}
	err := suite.cardRepository.AddCard(suite.clients[0].UserId, card)
	assert.Nil(suite.T(), err)

	foundCard, err := suite.cardRepository.GetCard("2card1")
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
		CvcCode: "000",
	}
	err := suite.cardRepository.UpdateCard(suite.clients[0].Cards[0].CardId, card)
	assert.Nil(suite.T(), err)

	foundCard, err := suite.cardRepository.GetCard(suite.clients[0].Cards[0].CardId)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), card.CvcCode, foundCard.CvcCode)
}

func (suite *CardRepositoryTestSuite) TestCardRepositoryDelete() {
	card := models.Card{
		CardId:   "2card2",
		Balance:  decimal.NewFromInt(100),
		ValidTo:  time.Now().Add(time.Hour),
		CvcCode:  "010",
		CardType: "debit",
		Currency: "RUB",
		ClientId: suite.clients[0].ClientId,
	}
	err := suite.cardRepository.AddCard(suite.clients[0].UserId, card)
	assert.Nil(suite.T(), err)

	err = suite.cardRepository.DeleteCard("2card2")
	assert.Nil(suite.T(), err)

	card, err = suite.cardRepository.GetCard("2card2")
	assert.ErrorIs(suite.T(), gorm.ErrRecordNotFound, err)
}

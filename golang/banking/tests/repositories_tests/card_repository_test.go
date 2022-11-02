package repositories_tests

import (
	"testing"

	"github.com/dopefresh/banking/golang/banking/src/di"
	"github.com/dopefresh/banking/golang/banking/src/models"
	"github.com/dopefresh/banking/golang/banking/src/repositories"
	"github.com/dopefresh/banking/golang/banking/src/utils"
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
	suite.setupData()
}

func (suite *CardRepositoryTestSuite) setupData() {
	suite.clientRepository.Db.Exec("CALL p_setup_test_data()")
}

func (suite *CardRepositoryTestSuite) TearDownSuite() {
	suite.clientRepository.Db.Exec(`SELECT truncate_tables('banking_db_user', '{"client","card","contract","credit","payment_schedule","transaction","user"}');`)
}

func (suite *CardRepositoryTestSuite) TestCardRepositoryCardCreate() {
}

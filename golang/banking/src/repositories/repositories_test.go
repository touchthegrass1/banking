package repositories

import (
	"testing"
	"time"

	"github.com/dopefresh/banking/golang/banking/src/database_layer"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	repository ClientRepository
	client     database_layer.Client
	user       database_layer.User
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) SetupSuite() {
	db := database_layer.InitializeDB()
	log := utils.ProvideLogger()
	suite.repository = ClientRepository{Db: db, Log: log}
	suite.setupData()
}

func (suite *RepositoryTestSuite) TearDownSuite() {
	suite.repository.Db.Exec(`SELECT truncate_tables('banking_db_user', '{"client","card","contract","credit","payment_schedule","transaction","user"}');`)
}

func (suite *RepositoryTestSuite) setupData() {
	user := database_layer.User{
		Password:    "example",
		LastLogin:   time.Now(),
		IsSuperuser: false,
		FirstName:   "Vasilii",
		LastName:    "Popov",
		Email:       "example@gmail.com",
		IsStaff:     false,
		IsActive:    true,
		DateJoined:  time.Now(),
		Phone:       "89999999999",
	}
	suite.repository.GetDB().Create(&user)

	client := database_layer.Client{
		RegistrationAddress: "-",
		ResidentialAddress:  "-",
		ClientType:          "individual",
		Ogrn:                "1201",
		Inn:                 "1110",
		Kpp:                 "1021",
		UserId:              user.Id,
	}
	suite.repository.GetDB().Create(&client)

	suite.user = user
	suite.client = client
}

func (suite *RepositoryTestSuite) TestClientRepositoryAddNewClient() {
	assert.NotNil(suite.T(), suite.repository)
	foundClient, err := suite.repository.GetClientByInn("1110")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.client, foundClient)
}

func (suite *RepositoryTestSuite) TestClientRepositoryUpdateClient() {
	client := suite.client
	client.ClientType = "jp"
	client.ResidentialAddress = "Moscow, Lva Tolstogo 1"
	client.RegistrationAddress = "Moscow, Lva Tolstogo 1"
	err := suite.repository.UpdateClient(client)
	assert.Nil(suite.T(), err)

	clientFound, err := suite.repository.GetClientByInn("1110")
	assert.Equal(suite.T(), clientFound, client)
}

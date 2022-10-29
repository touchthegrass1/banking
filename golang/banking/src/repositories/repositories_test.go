package repositories

import (
	"testing"

	"github.com/dopefresh/banking/golang/banking/src/database_layer"
	"github.com/dopefresh/banking/golang/banking/src/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	repository ClientRepository
}

func (suite *RepositoryTestSuite) SetupSuite() {
	db := database_layer.InitializeDB()
	log := utils.ProvideLogger()
	suite.repository = ClientRepository{Db: db, Log: log}
}

func (suite *RepositoryTestSuite) TearDownSuite() {
	suite.repository.Db.Exec("SELECT truncate_tables('banking_db_user');")
}

func (suite *RepositoryTestSuite) TestClientRepositoryAddNewClient() {
	assert.NotNil(suite.T(), suite.repository)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

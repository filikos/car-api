// +build integration

package Test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
	"workspace-go/coding-challange/car-api/api"
	"workspace-go/coding-challange/car-api/db"
	"workspace-go/coding-challange/car-api/model"

	"github.com/stretchr/testify/suite"
)

const (
	dbConfigTestPath = "../testdata/dbConfigTest.env"
	testDataPath     = "../testdata/testDataConnector.json"
)

type Suite struct {
	suite.Suite
	service api.Service
	cars    model.Cars
}

// Will run test suite
func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

// Will run before tests in the suite are run
func (suite *Suite) SetupSuite() {

	assert := suite.Assert()

	// Load testdata into suite
	_, err := os.Stat(testDataPath)
	os.IsNotExist(err)
	assert.Nil(err)

	file, err := ioutil.ReadFile(testDataPath)
	assert.Nil(err)

	err = json.Unmarshal([]byte(file), &suite.cars)
	assert.Nil(err)

	// Build and run API
	database, err := db.InitDB(dbConfigTestPath)
	assert.Nil(err)

	connector := &api.ConnectorDB{
		Database: *database,
	}

	suite.service = api.Service{
		Connector: connector,
	}
}

// Will run after all tests in the suite have been run
func (suite *Suite) TearDownSuite() {

	suite.service.Connector.CloseConnection()
}

// Will run before each test
func (suite *Suite) SetupTest() {

	assert := suite.Assert()

	database, err := db.InitDB(dbConfigTestPath)
	assert.Nil(err)

	defer database.Conn.Close()

	sqlStatement := `CREATE TABLE cars (
		id TEXT PRIMARY KEY,
		model TEXT,
		make TEXT,
		variant TEXT
	  );`

	_, err = database.Conn.Exec(sqlStatement)
	assert.Nil(err)

	// insert rows
	for _, c := range suite.cars {
		sqlStatement = `INSERT INTO cars (id, Model, Make, Variant)
		VALUES ($1, $2, $3, $4);`
		_, err = database.Conn.Exec(sqlStatement, c.ID, c.Model, c.Make, c.Variant)
		assert.Nil(err)
	}
}

// Will run after each test
func (suite *Suite) TearDownTest() {

	assert := suite.Assert()

	database, err := db.InitDB(dbConfigTestPath)
	assert.Nil(err)

	sqlStatement := `DROP TABLE cars`
	_, err = database.Conn.Exec(sqlStatement)
	assert.Nil(err)

	// table 'cars' should not exist at this point
	sqlStatement = `SELECT * FROM cars;`
	_, err = database.Conn.Exec(sqlStatement)
	assert.NotNil(err)

	database.Conn.Close()
}

func (suite *Suite) TestListCars() {

	got := suite.service.Connector.ListCars()
	want := suite.cars
	reflect.DeepEqual(want, got)
}

// +build integration

package Test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"workspace-go/coding-challange/car-api/api"
	"workspace-go/coding-challange/car-api/db"
	"workspace-go/coding-challange/car-api/model"

	"github.com/stretchr/testify/assert"
)

const (
	dbConfigTestPath = "../testdata/dbConfigTest.env"
)

func TestExample(t *testing.T) {

	database, err := db.InitDB(dbConfigTestPath)
	assert.Nil(t, err)

	var connector api.Controller
	connector = &api.ConnectorDB{
		Database: *database,
	}

	service := api.Service{
		Connector: connector,
	}

	req, err := http.NewRequest("GET", "/cars", nil)
	assert.Nil(t, err)

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ListCars)
	handler.ServeHTTP(respRec, req)

	want := http.StatusOK
	got := respRec.Code
	assert.Equal(t, got, want)

	var gotCars model.Cars
	err = json.NewDecoder(respRec.Body).Decode(&gotCars)
	assert.Nil(t, err)

	wantCars := service.Connector.ListCars()
	assert.Equal(t, gotCars, wantCars)
}

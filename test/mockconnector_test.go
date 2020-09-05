package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"workspace-go/coding-challange/car-api/api"
	"workspace-go/coding-challange/car-api/model"

	"github.com/stretchr/testify/assert"
)

func TestListCars(t *testing.T) {

	mockConnector := api.MockConnector{
		Data: model.Cars{
			{ID: "1", Model: "A45", Make: "mercedes", Variant: "amg"},
			{ID: "2", Model: "C", Make: "mercedes", Variant: "classic"},
			{ID: "3", Model: "B", Make: "mercedes", Variant: "casual"},
			{ID: "4", Model: "S", Make: "tesla", Variant: "sport"},
		},
	}

	service := api.Service{
		Connector: &mockConnector,
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

func TestCreateCarErrors(t *testing.T) {

	mockConnector := api.MockConnector{
		Data: model.Cars{
			{ID: "1", Model: "A45", Make: "mercedes", Variant: "amg"},
			{ID: "2", Model: "C", Make: "mercedes", Variant: "classic"},
			{ID: "3", Model: "B", Make: "mercedes", Variant: "casual"},
			{ID: "4", Model: "S", Make: "tesla", Variant: "sport"},
		},
	}

	service := api.Service{
		Connector: &mockConnector,
	}

	var tests = []struct {
		name       string
		statusWant int
		inputModel string
		inputMake  string
	}{
		{
			name:       "OK",
			statusWant: http.StatusOK,
			inputModel: "notEmpty",
			inputMake:  "notEmpty",
		},
		{
			name:       "BadRequest 1. Model empty",
			statusWant: http.StatusBadRequest,
			inputModel: "",
			inputMake:  "notEmpty",
		},
		{
			name:       "BadRequest 2 InputMake empty",
			statusWant: http.StatusBadRequest,
			inputModel: "notEmpty",
			inputMake:  "",
		},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {

			reqCar := model.Car{}
			reqCar.Model = test.inputModel
			reqCar.Make = test.inputMake

			b := new(bytes.Buffer)
			err := json.NewEncoder(b).Encode(reqCar)
			assert.Nil(t, err)

			req, err := http.NewRequest("POST", "/createCar", b)
			assert.Nil(t, err)

			respRec := httptest.NewRecorder()
			handler := http.HandlerFunc(service.CreateCar)
			handler.ServeHTTP(respRec, req)

			got := respRec.Code
			want := test.statusWant
			assert.Equal(t, got, want)
		})
	}
}

func TestCreateCar(t *testing.T) {

	mockConnector := api.MockConnector{
		Data: model.Cars{
			{ID: "1", Model: "A45", Make: "mercedes", Variant: "amg"},
			{ID: "2", Model: "C", Make: "mercedes", Variant: "classic"},
			{ID: "3", Model: "B", Make: "mercedes", Variant: "casual"},
			{ID: "4", Model: "S", Make: "tesla", Variant: "sport"},
		},
	}

	service := api.Service{
		Connector: &mockConnector,
	}

	wantCar := model.Car{
		ID:      "",
		Make:    "MyCar",
		Model:   "MyCarModel",
		Variant: "sport",
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(wantCar)
	assert.Nil(t, err)

	req, err := http.NewRequest("POST", "/createCar", b)
	assert.Nil(t, err)

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(service.CreateCar)
	handler.ServeHTTP(respRec, req)

	wantStatus := http.StatusOK
	gotStatus := respRec.Code
	assert.Equal(t, gotStatus, wantStatus)

	var gotCar model.Car
	err = json.NewDecoder(respRec.Body).Decode(&gotCar)
	assert.Nil(t, err)

	assert.Equal(t,  wantCar.Make, gotCar.Make)
	assert.Equal(t,  wantCar.Model, gotCar.Model)
	assert.Equal(t,  wantCar.Variant, gotCar.Variant)
	
}

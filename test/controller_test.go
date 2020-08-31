package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"workspace-go/coding-challange/car-api/api"
)

func TestGetCars(t *testing.T) {

	service := api.Service{
		CarData: []api.Car{
			{ID: "1", Model: "A45", Make: "mercedes", Variant: "amg"},
			{ID: "2", Model: "C", Make: "mercedes", Variant: "classic"},
			{ID: "3", Model: "B", Make: "mercedes", Variant: "casual"},
			{ID: "4", Model: "S", Make: "tesla", Variant: "sport"},
		},
	}

	req, err := http.NewRequest("GET", "/cars", nil)
	if err != nil {
		t.Fail()
	}

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ListCars)
	handler.ServeHTTP(respRec, req)

	want := http.StatusOK
	got := respRec.Code

	if want != got {
		t.Errorf("Expected Statuscode %v, got %v", want, got)
	}

	var respContent []api.Car
	err = json.NewDecoder(respRec.Body).Decode(&respContent)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(respContent, service.CarData) {
		t.Error("Response data does not equal service data")
	}
}

func TestCreateCarErrors(t *testing.T) {

	service := api.Service{
		CarData: []api.Car{
			{ID: "1", Model: "A45", Make: "mercedes", Variant: "amg"},
			{ID: "2", Model: "C", Make: "mercedes", Variant: "classic"},
			{ID: "3", Model: "B", Make: "mercedes", Variant: "casual"},
			{ID: "4", Model: "S", Make: "tesla", Variant: "sport"},
		},
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

			reqCar := api.Car{}
			reqCar.Model = test.inputModel
			reqCar.Make = test.inputMake

			b := new(bytes.Buffer)
			err := json.NewEncoder(b).Encode(reqCar)
			if err != nil {
				t.Error(err)
			}

			req, err := http.NewRequest("POST", "/createCar", b)
			if err != nil {
				t.Error(err)
			}

			respRec := httptest.NewRecorder()
			handler := http.HandlerFunc(service.CreateCar)
			handler.ServeHTTP(respRec, req)

		})
	}
}

func TestCreateCar(t *testing.T){

	service := api.Service{
		CarData: []api.Car{
			{ID: "1", Model: "A45", Make: "mercedes", Variant: "amg"},
			{ID: "2", Model: "C", Make: "mercedes", Variant: "classic"},
			{ID: "3", Model: "B", Make: "mercedes", Variant: "casual"},
			{ID: "4", Model: "S", Make: "tesla", Variant: "sport"},
		},
	}

	want := api.Car{
		ID: "",
		Make: "MyCar",
		Model: "MyCarModel",
		Variant: "sport",
	}

	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(want)
	if err != nil {
				t.Error(err)
	}

	req, err := http.NewRequest("POST", "/createCar", b)
	if err != nil {
		t.Fail()
	}

	respRec := httptest.NewRecorder()
	handler := http.HandlerFunc(service.CreateCar)
	handler.ServeHTTP(respRec, req)

	wantStatus := http.StatusOK
	got := respRec.Code

	if wantStatus != got {
		t.Errorf("Expected Statuscode %v, got %v", wantStatus, got)
	}
	
	var respContent api.Car
	err = json.NewDecoder(respRec.Body).Decode(&respContent)
	if err != nil {
		t.Error(err)
	}

	if want.Make != respContent.Make {
		t.Error("Property was not set")
	}

	if want.Model != respContent.Model {
		t.Error("Property was not set")
	}

	if want.Variant != respContent.Variant {
		t.Error("Property was not set")
	}
}
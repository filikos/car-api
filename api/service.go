package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"workspace-go/coding-challange/car-api/model"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	Connector Controller
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, msg string, err error) {

	log.Warn(fmt.Sprintf("Statuscode: %v, Message: %v, Error: %v", statusCode, msg, err))

	w.WriteHeader(statusCode)
	errResp := model.ErrorResponse{
		Code:    statusCode,
		Message: msg,
	}

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		log.Info(fmt.Sprintf("Unable to encode error response: %v", errResp))
	}
}

func (s *Service) CreateCar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newCar model.Car
	if err := json.NewDecoder(r.Body).Decode(&newCar); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Unable to read body. Body JSON format: {  'model' : 'string', 'make': 'string', 'variant' : 'string' }", err)
		return
	}

	if len(newCar.Make) == 0 || len(newCar.Model) == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "Model and Make are mandatory attributes.", nil)
	}

	newCar.ID = uuid.NewV4().String()
	car, err := s.Connector.AddCar(newCar)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Unable to createCar.", err)
		return
	}

	if err := json.NewEncoder(w).Encode(&car); err != nil {
		log.Info(fmt.Sprintf("CreateCar: Unable to encode %v", newCar))
	}
}

func (s *Service) ListCars(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s.Connector.ListCars()); err != nil {
		log.Info(fmt.Sprintf("Cars: Unable to encode cars %v", err))
	}
}

func (s *Service) GetCar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	reqID := params["id"]

	if _, err := uuid.FromString(reqID); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Unable to read ID. Please use UUID using RFC 4122 standard", err)
		return
	}

	car, err := s.Connector.GetCar(reqID)
	if err != nil {
		writeErrorResponse(w, http.StatusNotFound, "Unable to getCar.", err)
		return
	}

	if err := json.NewEncoder(w).Encode(*car); err != nil {
		log.Info(fmt.Sprintf("GetCar: Unable to encode car %v", err))
	}

}

func (s *Service) DeleteCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	reqID := params["id"]

	if _, err := uuid.FromString(reqID); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Unable to read ID. Please use UUID using RFC 4122 standard", err)
		return

	}

	err := s.Connector.DeleteCar(reqID)
	if err != nil {
		writeErrorResponse(w, http.StatusNotFound, err.Error(), err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Service) SearchByMake(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	makeValue := params["make"]

	if len(makeValue) == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "Attribute 'make' can`t be empty.", nil)
		return
	}

	cars := s.Connector.GetByMake(makeValue)
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Info(fmt.Printf("SearchByMake: Unable to encode %v", cars))
	}
}

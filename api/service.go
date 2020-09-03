package api

import (
	"encoding/json"
	"log"
	"net/http"
	"workspace-go/coding-challange/car-api/model"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Service struct {
	connector Controller
}



func writeErrorResponse(w http.ResponseWriter, statusCode int, msg string) {

	w.WriteHeader(statusCode)
	errResp := model.ErrorResponse{
		Code:    statusCode,
		Message: msg,
	}

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		log.Printf("Unable to encode error response: %v", errResp)
	}
}

func (s *Service) CreateCar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var car model.Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Unable to read body. Body JSON format: {  'model' : 'string', 'make': 'string', 'variant' : 'string' }")
		return
	}

	if len(car.Make) == 0 || len(car.Model) == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "Model and Make are mandatory attributes.")
	}

	car.ID = uuid.NewV4().String()
	s.connector.AddCar(car)

	if err := json.NewEncoder(w).Encode(&car); err != nil {
		log.Printf("CreateCar: Unable to encode %v", car)
	}
}

func (s *Service) ListCars(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s.connector.ListCars); err != nil {
		log.Printf("Cars: Unable to encode cars %v", err)
	}
}

func (s *Service) GetCar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	reqID := params["id"]

	if _, err := uuid.FromString(reqID); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Unable to read ID. Please use UUID using RFC 4122 standard")
		return
	}  

	car, err := s.connector.GetCar(reqID)
	if err != nil {
		log.Println(err)
		writeErrorResponse(w, http.StatusNotFound, err.Error())
	}

	if err := json.NewEncoder(w).Encode(*car); err != nil {
		log.Printf("GetCar: Unable to encode car %v", err)
	}
	
}

func (s *Service) DeleteCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	reqID := params["id"]

	if _, err := uuid.FromString(reqID); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Unable to read ID. Please use UUID using RFC 4122 standard")
		return

	} 
		
	err := s.connector.DeleteCar(reqID)
		if err != nil {
			log.Println(err)
			writeErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

	w.WriteHeader(http.StatusNoContent)	

}


func (s *Service) SearchByMake(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	makeValue := params["make"]

	if len(makeValue) == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "Attribute 'make' can`t be empty.")
		return
	}

	cars := s.connector.GetByMake(makeValue)
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Printf("SearchByMake: Unable to encode %v", cars)
	}
}

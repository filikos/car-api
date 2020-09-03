package api

import (
	"encoding/json"
	"fmt"
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
	s.CarData = append(s.CarData, car)
	s.connector.AddCar(car)

	if err := json.NewEncoder(w).Encode(&car); err != nil {
		log.Printf("CreateCar: Unable to encode %v", car)
	}
}

func (s *Service) ListCars(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s.CarData); err != nil {
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

	} else {

		for _, v := range s.CarData {
			if v.ID == reqID {
				if err := json.NewEncoder(w).Encode(v); err != nil {
					log.Printf("GetCar: Unable to encode car %v", err)
				}
				return
			}
		}
	}

	writeErrorResponse(w, http.StatusNotFound, fmt.Sprintf("Car with ID %v not found", reqID))
}

func (s *Service) DeleteCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	reqID := params["id"]

	if _, err := uuid.FromString(reqID); err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Unable to read ID. Please use UUID using RFC 4122 standard")
		return

	} else {
		for i, v := range s.CarData {
			if v.ID == reqID {
				s.CarData = append(s.CarData[:i], s.CarData[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		writeErrorResponse(w, http.StatusNotFound, fmt.Sprintf("Car with ID %v does not exist", reqID))
	}

}
func (s *Service) SearchByMake(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["make"]

	if len(name) == 0 {
		writeErrorResponse(w, http.StatusBadRequest, "Attribute: Make cant be empty.")
		return
	}

	cars := make([]model.Car, 0)
	for _, v := range s.CarData {
		if v.Make == name {
			cars = append(cars, v)
		}
	}

	if len(cars) > 0 {
		if err := json.NewEncoder(w).Encode(cars); err != nil {
			log.Printf("SearchByMake: Unable to encode %v", cars)
		}

		return
	}

	writeErrorResponse(w, http.StatusNotFound, fmt.Sprintf("Was not able to find car with name: %v.", name))
}

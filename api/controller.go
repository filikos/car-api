package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Service struct {
	CarData []Car
}

func (s *Service) CreateCar(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var car Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResp := ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Unable to read body. Body JSON format: {  'model' : 'string', 'make': 'string', 'variant' : 'string' }",
		}

		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			log.Printf("CreateCar: Unable to encode %v", errResp)
		}
		return
	}

	ID := uuid.NewV4()
	car.ID = ID.String()

	s.CarData = append(s.CarData, car)
	if err = json.NewEncoder(w).Encode(&car); err != nil {
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

	_, err := uuid.FromString(reqID)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		errResp := ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Unable to read ID. Please use UUID using RFC 4122 standard",
		}

		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			log.Printf("CreateCar: Unable to encode %v", errResp)
		}

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

	w.WriteHeader(http.StatusNotFound)
	errResp := ErrorResponse{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("Car with ID %v not found", reqID),
	}

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		log.Printf("GetCar: Unable to encode %v", errResp)
	}
}

func (s *Service) DeleteCar(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	reqID := params["id"]

	_, err := uuid.FromString(reqID)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		errResp := ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Unable to read ID. Please use UUID using RFC 4122 standard",
		}

		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			log.Printf("CreateCar: Unable to encode %v", errResp)
		}

		return

	} else {
		for i, v := range s.CarData {
			if v.ID == reqID {
				s.CarData = append(s.CarData[:i], s.CarData[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		errResp := ErrorResponse{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("Car with ID %v does not exist", reqID),
		}

		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			log.Printf("CreateCar: Unable to encode %v", errResp)
		}

	}

}
func (s *Service) SearchByMake(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["name"]

	if len(name) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		errResp := ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Make cant be empty.",
		}

		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			log.Printf("SearchByMake: Unable to encode %v", errResp)
		}
		return
	}

	cars := make([]Car, 0)
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

	w.WriteHeader(http.StatusNotFound)
	errResp := ErrorResponse{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("Was not able to find car with name: %v.", name),
	}

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		log.Printf("SearchByMake: Unable to encode %v", errResp)
	}
}

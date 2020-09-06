package api

import (
	"workspace-go/coding-challange/car-api/model"
)
/*
	This connector is mocking the database by storing all information in-memory.
	If there are 3rd party services which are not suitable for development purposes they can be mocked here.
*/
type MockConnector struct {
	Data model.Cars
}

func (mock *MockConnector) CloseConnection() {
}

func (mock *MockConnector) AddCar(newCar model.Car) (*model.Car, error) {
	mock.Data.AddCar(newCar)
	return &newCar, nil
}

func (mock *MockConnector) GetCar(ID string) (*model.Car, error) {
	return mock.Data.GetCar(ID)
}

func (mock *MockConnector) DeleteCar(ID string) error {
	return mock.Data.Delete(ID)
}

func (mock *MockConnector) ListCars() model.Cars {
	return mock.Data
}

func (mock *MockConnector) GetByMake(makeValue string) model.Cars {
	return mock.Data.GetByMake(makeValue)
}

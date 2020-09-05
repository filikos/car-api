package api

import (
	"workspace-go/coding-challange/car-api/model"
)

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

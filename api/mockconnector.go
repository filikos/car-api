package api

import "workspace-go/coding-challange/car-api/model"

type MockConnector struct {
	data *model.Cars
}

func (mock *MockConnector) AddCar(newCar model.Car) (*model.Car, error) {

	mock.data.AddCar(newCar)
	return &newCar, nil
}

func (mock *MockConnector) GetCar(ID string) (*model.Car, error) {

	return mock.data.GetCar(ID)
}

func (mock *MockConnector) DeleteCar(ID string) error {

	return mock.data.Delete(ID)
}

func (mock *MockConnector) ListCars() model.Cars {

	return *mock.data
}

func (mock *MockConnector) GetByMake(makeValue string) model.Cars {

	return mock.data.GetByMake(makeValue)
}

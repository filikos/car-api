package api

import (
	"workspace-go/coding-challange/car-api/db"
	"workspace-go/coding-challange/car-api/model"
)

type Controller interface {
	AddCar(newCar model.Car) (*model.Car, error)
	GetCar(ID string) (*model.Car, error)
	DeleteCar(ID string) error
	ListCars() model.Cars
	GetByMake(makeValue string) model.Cars
}

type ConnectorDB struct {
	Database db.Database
}

func (c *ConnectorDB) AddCar(newCar model.Car) (*model.Car, error) {
	return &model.Car{}, nil
}
func (c *ConnectorDB) GetCar(ID string) (*model.Car, error) {
	return &model.Car{}, nil
}
func (c *ConnectorDB) DeleteCar(ID string) error {
	return nil
}
func (c *ConnectorDB) ListCars() model.Cars {
	return model.Cars{}
}
func (c *ConnectorDB) GetByMake(makeValue string) model.Cars {
	return model.Cars{}
}



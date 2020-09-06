package api

import (
	"fmt"
	"workspace-go/coding-challange/car-api/db"
	"workspace-go/coding-challange/car-api/model"
	log "github.com/sirupsen/logrus"
)

// Interface providing all functions needed to fullfill the services needs.
type Controller interface {
	CloseConnection() 
	AddCar(newCar model.Car) (*model.Car, error)
	GetCar(ID string) (*model.Car, error)
	DeleteCar(ID string) error
	ListCars() model.Cars
	GetByMake(makeValue string) model.Cars
}

// Connector containing all informations for accessing the database and calls the database functions.
type ConnectorDB struct {
	Database db.Database
}

func (c *ConnectorDB) CloseConnection()  {

	if err := c.Database.Conn.Close(); err != nil {
		log.Warn(fmt.Sprintf("Failed to close DB connection: %v",err))
	}
}

func (c *ConnectorDB) AddCar(newCar model.Car) (*model.Car, error) {
	return c.Database.AddCar(newCar)
}
func (c *ConnectorDB) GetCar(ID string) (*model.Car, error) {
	return c.Database.GetCar(ID)
}
func (c *ConnectorDB) DeleteCar(ID string) error {
	return c.Database.DeleteCar(ID)
}
func (c *ConnectorDB) ListCars() model.Cars {
	return c.Database.ListCars()
}
func (c *ConnectorDB) GetByMake(makeValue string) model.Cars {
	return c.Database.GetByMake(makeValue)
}

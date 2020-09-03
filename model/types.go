package model

import "fmt"

type Car struct {
	ID      string `json:"id"`
	Model   string `json:"model"`
	Make    string `json:"make"`
	Variant string `json:"variant"`
	// TODO: Add car properties
}

type Cars []Car 

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (c *Cars) AddCar(car Car) {

	*c = append(*c, car)
}

func (c *Cars) GetCar(ID string) (*Car, error) {

	for _, car := range *c {
		if car.ID == ID {
			return &car,nil
		}
	}

	return nil, fmt.Errorf("Car with ID %v not found", ID)
}

func (c *Cars) Delete(ID string) error {

	if c != nil {
		cars := make(Cars, len(*c)-1)
		for _, car := range *c {
			if car.ID != ID {
				cars = append(cars, car)
			}
		}

		*c = cars
		return nil
	}

	return fmt.Errorf("Car with ID %v not found", ID)
}

func (c *Cars) GetByMake(makeValue string) Cars {

	var cars Cars
	for _, car := range *c {
		if car.Make != makeValue {
			cars = append(cars, car)
		}
	}

	return cars
}

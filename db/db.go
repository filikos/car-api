package db

import (
	"database/sql"
	"fmt"
	"time"

	"os"
	"workspace-go/coding-challange/car-api/model"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

const (
	connTimeoutSec = 5

	dbConnectionAttemts = 5
	retryInterval       = 2
)

// Provides functionality for initializing and execute database operations.
type Database struct {
	Conn *sql.DB
}

func InitDB(configPath string) (*Database, error) {

	err := godotenv.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to load DB configuration %v", err)
	}

	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")

	var dbPool *sql.DB
	url := fmt.Sprintf("user=%v dbname=%v host=%v port=%v password=%v  connect_timeout=%v sslmode=disable", dbUser, dbName, dbHost, dbPort, dbPassword, connTimeoutSec)

	for i := 0; i < dbConnectionAttemts; i++ {

		dbPool, err = sql.Open("postgres", url)
		if err != nil {
			return nil, err
		}

		err = dbPool.Ping()
		if err == nil {

			dbPool.SetMaxOpenConns(7)
			dbPool.SetMaxIdleConns(5)

			log.Info("Database connection established")
			return &Database{dbPool}, nil
		}

		log.Warn(fmt.Sprintf("Connecting to DB try: %v", i+1))
		time.Sleep(retryInterval * time.Second)
	}

	return nil, fmt.Errorf("Failed to connect database. Server will be shut down. Error: %v", err)
}

func (db *Database) AddCar(newCar model.Car) (*model.Car, error) {

	sqlStatement :=
		`INSERT INTO cars (id, model, make, variant)	
	VALUES ($1, $2, $3, $4)`

	_, err := db.Conn.Query(sqlStatement, newCar.ID, newCar.Model, newCar.Make, newCar.Variant)
	if err != nil {
		return nil, fmt.Errorf("Query: %v", err)
	}

	return &newCar, nil
}

func (db *Database) GetCar(ID string) (*model.Car, error) {

	var car model.Car
	sqlStatement := `SELECT * FROM cars WHERE ID = $1`
	err := db.Conn.QueryRow(sqlStatement, ID).Scan(&car.ID, &car.Model, &car.Make, &car.Variant)
	if err != nil {
		return nil, fmt.Errorf("Query: %v", err)
	}

	return &car, nil
}

func (db *Database) DeleteCar(ID string) error {

	sqlStatement := `DELETE FROM cars WHERE ID = $1`
	_, err := db.Conn.Exec(sqlStatement, ID)
	if err != nil {
		return fmt.Errorf("Query: %v", err)
	}

	return nil
}

func (db *Database) ListCars() model.Cars {

	rows, err := db.Conn.Query("SELECT * FROM cars")
	if err != nil {
		log.Warnf("Query: Unable to listCars: %v", err)
		return nil
	}

	defer rows.Close()

	var cars model.Cars
	for rows.Next() {

		var car model.Car
		err := rows.Scan(&car.ID, &car.Model, &car.Make, &car.Variant)
		if err != nil {
			log.Info(fmt.Sprintf("Unable to scan row:%v", err))
			continue
		}

		cars = append(cars, car)
	}

	return cars
}

func (db *Database) GetByMake(makeValue string) model.Cars {

	rows, err := db.Conn.Query("SELECT * FROM cars WHERE make = $1", makeValue)
	if err != nil {
		log.Warnf("Query: Unable to GetByMake: %v", err)
		return nil
	}

	defer rows.Close()

	var cars model.Cars
	for rows.Next() {

		var car model.Car
		err := rows.Scan(&car.ID, &car.Model, &car.Make, &car.Variant)
		if err != nil {
			log.Info(fmt.Sprintf("Unable to scan row: %v", err))
			continue
		}

		cars = append(cars, car)
	}

	return cars
}

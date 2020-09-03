package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	HOST = "database"
	PORT = 5432
)

type Database struct {
	Conn *sql.DB
}

func InitDB(configPath string) (*Database, error) {

	err := godotenv.Load(configPath)
	if err != nil {
		return nil, fmt.Errorf("Unable to load DB configuration %v", err)
	}

	dbUser := os.Getenv("dbUser")
	dbPassword := os.Getenv("dbPassword")
	dbName := os.Getenv("dbName")

	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, dbUser, dbPassword, dbName)
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	db := &Database{}
	db.Conn = conn

	err = db.Conn.Ping()
	if err != nil {
		return nil, err
	}

	log.Printf("Database successfully connected")
	return db, nil
}

// TODO: Add DB operations for endpoints
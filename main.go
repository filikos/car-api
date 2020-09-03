package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"workspace-go/coding-challange/car-api/api"

	"workspace-go/coding-challange/car-api/model"
	"workspace-go/coding-challange/car-api/db"

	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:    "Car-Management-API",
		Version: "v0.0.0",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Port the Rest-API will listen on.",
				Value: 8080,
			},
			&cli.BoolFlag{
				Name:        "mockmode",
				Usage:       "Set 'true' to use mocked mode.",
				Value:       false,
				DefaultText: "API will use DB connection",
			},
		},
	}

	port := 8080
	var connector api.Controller
	app.Action = func(c *cli.Context) error {

		port = c.Int("port")
		if c.Bool("mockmode") {
			connector = &api.MockConnector{
				Data: model.Cars{
					{ID: "1", Model: "A45", Make: "mercedes", Variant: "amg"},
					{ID: "2", Model: "C", Make: "mercedes", Variant: "classic"},
					{ID: "3", Model: "B", Make: "mercedes", Variant: "casual"},
					{ID: "4", Model: "S", Make: "tesla", Variant: "sport"},
					{ID: "5", Model: "3", Make: "tesla", Variant: "tour"},
					{ID: "6", Model: "X", Make: "tesla", Variant: "midnight"},
					{ID: "7", Model: "Y", Make: "tesla", Variant: "standart"},
				},
			}
		} else {
			db, err := db.InitDB("./config/dbConfig.env")
			if err != nil {
				log.Printf("Failed to connect database. Server will be shut down. Error: %v", err)
				os.Exit(0)
			}

			connector = &api.ConnectorDB{
				Database: *db,
			}
		}

		service := api.Service{
			Connector: connector,
		}

		startServer(service, port)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func startServer(service api.Service, port int) {

	r := mux.NewRouter()
	// TODO: Host to be set here

	r.HandleFunc("/createCar", service.CreateCar).Methods("POST")
	r.HandleFunc("/cars", service.ListCars).Methods("GET")
	r.HandleFunc("/cars/{id}", service.GetCar).Methods("GET")
	r.HandleFunc("/cars/{id}", service.DeleteCar).Methods("DELETE")
	r.HandleFunc("/search/{make}", service.SearchByMake).Methods("GET")

	server := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%v", port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Printf("Server Running on port: %v", port)
	<-done
	log.Println("Server Stopped")
	log.Println("shutting down...")
	os.Exit(0)
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"workspace-go/coding-challange/car-api/api"
	"workspace-go/coding-challange/car-api/db"

	"workspace-go/coding-challange/car-api/model"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const(
	dbConnectionAttemts = 5
	retryInterval = 2
)

func main() {

	// Set up CLI application with all available flags.
	app := &cli.App{
		Name:    "Car-Management-API",
		Version: "v1.0.0",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Port the Rest-API will listen on.",
				Value: 8080,
			},
			&cli.PathFlag{
				Name:        "configPath",
				Usage:       "Path to *.env postgres config file.",
				Value:       "./config/dbConfig.env",
				DefaultText: "./config/dbConfig.env",
			},
			&cli.BoolFlag{
				Name:        "mockmode",
				Usage:       "Set 'true' to use mocked mode.",
				Value:       false,
				DefaultText: "API will use DB connection",
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Usage:       "Set 'true' to enable verbose DEBUG-level logging.",
				Value:       false,
				DefaultText: "Logging on WARN-level",
			},
		},
	}


	log.SetFormatter(&log.JSONFormatter{})
	var connector api.Controller
	
	// Reading the CLI-Arguments, build and start the service
	app.Action = func(c *cli.Context) error {

		if c.Bool("verbose") {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.WarnLevel)
		}

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

			var database *db.Database
			var err error
			for i := 0; i < dbConnectionAttemts; i++ {
				
				log.Warn(fmt.Sprintf("Connecting to DB try: %v", i+1))
				database, err = db.InitDB(c.Path("configPath"))
				if err == nil {
					break
				}

				time.Sleep(retryInterval * time.Second)
			}

			if err != nil {
				log.Errorf("Failed to connect database. Server will be shut down. Error: %v", err)
				os.Exit(0)
			}	

			log.Info("Database connection established.")
			connector = &api.ConnectorDB{
				Database: *database,
			}
		}

		service := api.Service{
			Connector: connector,
		}

		startServer(service, c.Int("port"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

// Setting up routes with handlers and start the server. Will wait for termination signal to perform gracefull shut down.
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
			log.Fatal(err)
		}
	}()

	log.Info(fmt.Sprintf("Server Running on port: %v", port))
	defer service.Connector.CloseConnection()
	<-done
	log.Info("Server Stopped")
}

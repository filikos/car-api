package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"workspace-go/coding-challange/car-api/api"


	"github.com/gorilla/mux"
)

func main() {

	api.CarData = append(api.CarData, api.Car{ID: "1", Model: "amg", Make: "mercedes", Variant: "sports"})

	r := mux.NewRouter()
	// TODO: Host to be set here

	r.HandleFunc("/createCar", api.CreateCar).Methods("POST")
	r.HandleFunc("/cars", api.Cars).Methods("GET")
	r.HandleFunc("/cars/{id}", api.GetCar).Methods("GET")
	r.HandleFunc("/cars/{id}", api.DeleteCar).Methods("DELETE")
	r.HandleFunc("/search/{name}", api.SearchByMake).Methods("GET")

	server := &http.Server{
		Handler: r, // Pass our instance of gorilla/mux in.
		Addr:    "127.0.0.1:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
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

	
	log.Print("Server Running")
	<-done
	log.Print("Server Stopped Running")
	log.Println("shutting down")
	os.Exit(0)
}

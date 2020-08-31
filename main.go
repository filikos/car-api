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

	service := api.Service{
		CarData: []api.Car{
			{ID: "1", Model: "A45", Make: "mercedes", Variant: "amg"},
			{ID: "2", Model: "C", Make: "mercedes", Variant: "classic"},
			{ID: "3", Model: "B", Make: "mercedes", Variant: "casual"},
			{ID: "4", Model: "S", Make: "tesla", Variant: "sport"},
			{ID: "5", Model: "3", Make: "tesla", Variant: "tour"},
			{ID: "6", Model: "X", Make: "tesla", Variant: "midnight"},
			{ID: "7", Model: "Y", Make: "tesla", Variant: "standart"},
		},
	}

	r := mux.NewRouter()
	// TODO: Host to be set here

	r.HandleFunc("/createCar", service.CreateCar).Methods("POST")
	r.HandleFunc("/cars", service.ListCars).Methods("GET")
	r.HandleFunc("/cars/{id}", service.GetCar).Methods("GET")
	r.HandleFunc("/cars/{id}", service.DeleteCar).Methods("DELETE")
	r.HandleFunc("/search/{name}", service.SearchByMake).Methods("GET")

	server := &http.Server{
		Handler: r, 
		Addr:    "127.0.0.1:8080",
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

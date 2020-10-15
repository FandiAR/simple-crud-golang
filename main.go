package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FandiAR/simple-crud-golang/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	endPoint := fmt.Sprintf(":%d", 9000)
	maxHeaderBytes := 1 << 20
	personHandler := handler.Person{}

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    5000,
		WriteTimeout:   5000,
		MaxHeaderBytes: maxHeaderBytes,
	}

	router.HandleFunc("/people", personHandler.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", personHandler.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people", personHandler.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", personHandler.DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(server.ListenAndServe())
}

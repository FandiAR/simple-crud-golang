package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FandiAR/simple-crud-golang/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	endPoint := fmt.Sprintf(":%d", 8080)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    5000,
		WriteTimeout:   5000,
		MaxHeaderBytes: maxHeaderBytes,
	}

	router.HandleFunc("/people", controllers.GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", controllers.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", controllers.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", controllers.DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(server.ListenAndServe())
}

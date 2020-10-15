package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FandiAR/simple-crud-golang/models"

	"github.com/gorilla/mux"
)

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	people := models.PersonData()
	params := mux.Vars(req)

	itemId, _ := strconv.Atoi(params["id"])
	for _, item := range people {
		if item.ID == itemId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	people := models.PersonData()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	people := models.PersonData()
	params := mux.Vars(req)
	var person models.Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID, _ = strconv.Atoi(params["id"])
	people = append(people, person)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	people := models.PersonData()
	params := mux.Vars(req)
	itemId, _ := strconv.Atoi(params["id"])
	for index, item := range people {
		if item.ID == itemId {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

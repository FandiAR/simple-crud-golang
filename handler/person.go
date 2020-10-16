package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FandiAR/simple-crud-golang/models"
	"github.com/FandiAR/simple-crud-golang/repository"

	"github.com/gorilla/mux"
)

type Person struct {
	personRepo *repository.Person
}

func (p *Person) GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	itemId, err := strconv.ParseInt(string(params["id"]), 0, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error happened"))
		return
	}
	person, err := p.personRepo.GetById(itemId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Person not found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func (p *Person) GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	people, err := p.personRepo.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

func (p *Person) CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	var person models.Person
	err := json.NewDecoder(req.Body).Decode(&person)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = p.personRepo.Create(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Person created"))
}

func (p *Person) DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	itemId, err := strconv.ParseInt(string(params["id"]), 0, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error happened"))
		return
	}
	err = p.personRepo.Delete(itemId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error happened"))
		return
	}
	w.Write([]byte("Person deleted"))
}

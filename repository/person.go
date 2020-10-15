package repository

import (
	"errors"

	"github.com/FandiAR/simple-crud-golang/staticdata"

	"github.com/FandiAR/simple-crud-golang/models"
)

type Person struct {
}

func (p *Person) GetAll() (people []models.Person, err error) {
	people = staticdata.GetPeople()
	if len(people) <= 0 {
		err = errors.New("Not found")
	}
	return
}

func (p *Person) GetById(id int64) (person models.Person, err error) {
	people := staticdata.GetPeople()
	for _, item := range people {
		if int64(item.ID) == id {
			person = item
			return
		}
	}
	err = errors.New("Person not found")
	return
}

func (p *Person) Create(person models.Person) (err error) {
	size := len(staticdata.GetPeople())
	person.ID = size + 1
	staticdata.AddPerson(person)
	return
}

func (p *Person) Delete(id int64) (err error) {
	staticdata.RemovePerson(id)
	return
}

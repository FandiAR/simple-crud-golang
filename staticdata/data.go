package staticdata

import "github.com/FandiAR/simple-crud-golang/models"

var people []models.Person

func init() {
	people = append(people, models.Person{ID: 1, Firstname: "Nic", Lastname: "Raboy", Address: &models.Address{City: "Dublin", State: "CA"}})
	people = append(people, models.Person{ID: 2, Firstname: "Maria", Lastname: "Raboy"})
}

func GetPeople() []models.Person {
	return people
}

func AddPerson(person models.Person) {
	people = append(people, person)
}

func RemovePerson(id int64) {
	for index, item := range people {
		if int64(item.ID) == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
}

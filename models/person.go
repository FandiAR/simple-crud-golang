package models

type Person struct {
	ID        int      `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

func PersonData() []Person {
	var people []Person
	people = append(people, Person{ID: 1, Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
	people = append(people, Person{ID: 2, Firstname: "Maria", Lastname: "Raboy"})
	return people
}

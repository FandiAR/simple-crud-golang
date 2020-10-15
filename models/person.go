package models

type Person struct {
	ID        int      `json:"id"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

func PersonData() []Person {
	var people []Person
	people = append(people, Person{ID: 1, Firstname: "Nic", Lastname: "Raboy", Address: &Address{City: "Dublin", State: "CA"}})
	people = append(people, Person{ID: 2, Firstname: "Maria", Lastname: "Raboy"})
	return people
}

package models

import (
	"log"
)

type Person struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

// Cart .
type Cart struct {
	ID       string    `json:"id,omitempty"`
	Articles []Article `json:"articles,omitempty"`
}

// Article .
type Article struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Price string `json:"price,omitempty"`
}

func populatePerson(people Person) {
	log.Println("populatePerson")
	/*
		people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
		people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	*/
}

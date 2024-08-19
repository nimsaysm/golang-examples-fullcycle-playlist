package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string
	Email string
	CPF   int
}

type InternationalPerson struct {
	Person
	Country string `json: country`
}

// bind func -> is a method from Car
func (p Person) walk() {
	fmt.Println(p.Name, "walked")
}

func (i InternationalPerson) presentation() {
	fmt.Printf("Name: %s, Email: %s, CPF: %d, Country: %s/n", i.Name, i.Email, i.CPF, i.Country)
}

func main() {
	person := Person{
		Name: "Rubia",
	}

	person.walk()

	person2 := InternationalPerson{
		Person: Person{
			Name:  "Yasmin",
			Email: "yasmin@gmail.com",
			CPF:   11111111111,
		},
		Country: "Canada",
	}

	person2.presentation()

	//Marshal -> transforms Struct in Json
	personJson, err := json.Marshal(person2)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(personJson))

	//Json to string
	jsonPerson3 := `{"Name":"Yasmin","Email":"yasmin@gmail.com","CPF":11111111111,"Country":"Canada"}`
	person3 := InternationalPerson{}

	json.Unmarshal([]byte(jsonPerson3), &person3)
	fmt.Println(person3)
}

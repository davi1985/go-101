package main

import "fmt"

type Person struct {
	Name   string
	Age    uint8
	Status bool
}

func (p Person) String() string {
	return fmt.Sprintf("I'm %s, I'm %d", p.Name, p.Age)
}

type NaturalPerson struct {
	Person
	lastName string
	cpf      string
}

type LegalEntity struct {
	cnpj          string
	corporateName string
}

func main() {
	person1 := NaturalPerson{
		Person:   Person{Name: "Davi", Age: 39, Status: true},
		lastName: "Silva",
		cpf:      "001.002.003-44",
	}

	fmt.Println(person1)
}

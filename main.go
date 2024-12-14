package main

import "fmt"

type Person struct {
	Name       string
	Age        uint8
	Occupation string
}

func main() {
	davi := Person{
		Name:       "Davi Silva",
		Occupation: "Software Engineer",
		Age:        39,
	}

	joelma := Person{"Joelma", 36, "Wife"}

	fmt.Println(davi)
	fmt.Println(joelma)
}

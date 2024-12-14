package main

import "fmt"

var (
	name string
	age  int
)

func main() {
	var (
		fullname   string
		occupation string
	)

	fullname = "Davi Silva"
	occupation = "Software Engineer"

	message := "Go 101 - lesson 03"

	fmt.Println(message)
	fmt.Println(fullname)
	fmt.Println(occupation)
}

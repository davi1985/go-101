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
	// var a, b, c int32             // multiples declaration
	// var d, e, f = true, 2.3, "Hi" // multiples attribution

	fmt.Println(message)
	fmt.Println(fullname)
	fmt.Println(occupation)

	// swap values
	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}

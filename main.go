package main

import "fmt"

var (
	heads, tails int
)

func throwCoin(side string) {
	switch side {
	case "heads":
		heads++
	case "tails":
		tails++

	default:
		fmt.Println("try again")
	}
}

func main() {

	throwCoin("heads")

	fmt.Println(heads)
	fmt.Println(tails)
}

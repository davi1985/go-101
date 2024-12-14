package main

import "fmt"

func main() {
	names := []string{"Davi", "Joelma", "Laura", "Sarah"}
	var i int

	for i < len(names) {
		fmt.Println(names[i])
		i++
	}
}

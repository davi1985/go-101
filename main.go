package main

import "fmt"

func main() {
	names := []string{"Davi", "Joelma", "Laura", "Sarah"}

	names = append(names, "Bony")

	fmt.Println(names, len(names), cap(names))
}

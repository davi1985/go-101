package main

import "fmt"

func main() {
	sayHello("Davi Silva")
	fmt.Println("Total:", sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}

func sayHello(name string) {
	fmt.Println(name)
}

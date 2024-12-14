package main

import (
	"fmt"
	"strconv"
)

func sum(a, b int) int {
	return a + b
}

func sayHello(name string) {
	fmt.Println(name)
}

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)

	if err != nil {
		return
	}

	total = a + i
	return
}

func main() {
	sayHello("Davi Silva")
	fmt.Println("Total:", sum(1, 2))

	total, err := convertAndSum(1, "5")

	if err != nil {
		fmt.Println((err))
	}

	fmt.Println("Total:", total)
}

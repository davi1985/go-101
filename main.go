package main

import "fmt"

func main() {
	ages := make(map[string]uint8)
	ages["davi"] = 39
	ages["joelma"] = 34
	ages["laura"] = 2

	fmt.Println(ages)

	val, ok := ages["davi"]
	fmt.Println(val, ok)
	value, has := ages["bony"]
	fmt.Println(value, has)
}

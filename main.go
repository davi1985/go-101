package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	a, b := 10, 20

	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a equals b")
	}

	file, err := os.Open("hello.txt")

	if err != nil {
		log.Panic(err)

		return
	}

	data := make([]byte, 100)

	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}

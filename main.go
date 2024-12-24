package main

import (
	"fmt"
	"time"
)

func numbers(done chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}

	done <- true
}

func letters(done chan<- bool) {
	for i := 'a'; i < 'l'; i++ {
		fmt.Printf("%c ", i)
		time.Sleep(time.Millisecond * 230)
	}

	done <- true
}

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go numbers(ch1)
	go letters(ch2)

	<-ch1
	<-ch2

	fmt.Println("End")
}

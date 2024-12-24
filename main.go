package main

import (
	"fmt"
	"time"
)

func numbers(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("[Write by channel: %d ]", i)
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Println("Write done")
	close(n)
}

func main() {
	ch1 := make(chan int, 3)

	go numbers(ch1)

	for v := range ch1 {
		fmt.Printf("read by channel: %d\n", v)
		time.Sleep(time.Millisecond * 150)
	}

	fmt.Println("End")
}

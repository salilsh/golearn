package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 100)

	go sendmultiple(c)

	for v := range c {
		fmt.Println(v)

	}

	fmt.Println("exiting")

}

func sendmultiple(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

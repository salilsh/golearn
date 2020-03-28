package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	fmt.Println("Launching go routine")

	go writeValues(ch)

	fmt.Println("already out of go-routine, into main...")

	readValues(ch)

	fmt.Println("Exiting...")

}

func writeValues(c chan<- int) {
	fmt.Println("entered go routine")
	for i := 1; i <= 100; i++ {
		fmt.Println("wrote value", i, "to channel")
		c <- i
	}
	fmt.Println("Oh, here's the close signal!")
	close(c)

}

func readValues(c <-chan int) {
	fmt.Println("reading from channel...")
	for {
		select {
		case v, ok := <-c:
			{
				if ok {
					fmt.Println("haven't hit close, yet...reading", v, "from channel")
					//fmt.Println(v)
				} else {
					return
				}
			}
		}
	}
}

// alternative solution

/* func readValues(c <-chan int) {
	for {
		v, ok := <-c
		if ok {
			fmt.Println(v)

		} else {
			return
		}

	}

} */

//Todd's solution (simplest and best)

/* func readValues(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
} */

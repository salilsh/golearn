package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)

	for k := 1; k < 11; k++ {
		go writeToChannel(ch, ((k * 10) + 1))

	}

	readValues(ch)

	close(ch)

	fmt.Println("Exiting....")
}

func writeToChannel(c chan<- int, end int) {
	for i := end - 10; i < end; i++ {
		fmt.Println("writing value", i, "to channel")
		c <- i
		//fmt.Println("ROUTINES", runtime.NumGoroutine())
	}

}

func readValues(c <-chan int) {

	for i := 1; i < 101; i++ {

		//fmt.Println("reading value", <-c, "from channel")
		fmt.Println(i, <-c)
	}

}

/* func readValues(c <-chan int) {
	i := 1
	for {
		select {
		case v, ok := <-c:
			{
				if ok {
					fmt.Println(i, v)
					i++
				} else {
					return
				}
			}
		}
	}
}*/

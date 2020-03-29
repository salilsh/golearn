package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	ch := make(chan int)

	//implementing by looking at fellow student's work. lesson: if a channel is closed before values are even received from it, it leads to a deadlock.
	go func() {
		for k := 1; k < 11; k++ {
			wg.Add(1)
			go writeToChannel(ch, ((k * 10) + 1))

		}
		wg.Wait()
		close(ch)
	}()

	readValues(ch)

	fmt.Println("Exiting....")
}

func writeToChannel(c chan<- int, end int) {
	for i := end - 10; i < end; i++ {
		//fmt.Println("writing value", i, "to channel")
		c <- i

		//fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	wg.Done()

}

/* func readValues(c <-chan int) {

	for i := 1; i < 101; i++ {

		//fmt.Println("reading value", <-c, "from channel")
		fmt.Println(i, <-c)
	}

} */

func readValues(c <-chan int) {
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
}

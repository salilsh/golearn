package main

import (
	"fmt"
	"sync"
)

func main() {
	odd := make(chan int)
	even := make(chan int)

	receive := make(chan int)

	go compute(odd, even)

	go fanIn(odd, even, receive)

	for v := range receive {
		fmt.Println(v)
	}

}

func compute(o, e chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func fanIn(o, e <-chan int, r chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range e {
			r <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			r <- v
		}
		wg.Done()
	}()
	wg.Wait()

	close(r)
}

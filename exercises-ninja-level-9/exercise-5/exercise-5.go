// this produces an error if it's run with go run -race file.go
// since multiple goroutines write to the same variable that get
// executed at a different time

// fixed using atomic

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64
	var wg sync.WaitGroup

	//counter = 0

	gs := 100

	wg.Add(gs)

	fmt.Println(runtime.NumGoroutine())

	for i := 0; i < 100; i++ {

		fmt.Println("go routines rn", runtime.NumGoroutine())
		fmt.Println("iter", i)
		go func() {
			atomic.AddInt64(&counter, 1)
			r := atomic.LoadInt64(&counter)
			fmt.Println(r)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("we're here", counter)
}

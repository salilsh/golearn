// this produces an error if it's run with go run -race file.go
// since multiple goroutines write to the same variable that get
// executed at a different time

//fixed using mutex

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup

	wg.Add(100)

	var m sync.Mutex

	fmt.Println(runtime.NumGoroutine())

	for i := 0; i < 100; i++ {
		fmt.Println("iter", i)
		go func() {
			m.Lock()
			v := counter
			runtime.Gosched()
			v++
			fmt.Println("NGR", runtime.NumGoroutine())
			counter = v
			m.Unlock()
			fmt.Println("in-loop ctr", counter)
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("we're here", counter)
}

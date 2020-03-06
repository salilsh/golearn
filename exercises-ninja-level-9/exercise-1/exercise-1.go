package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("this is from the main go routine")
	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println(runtime.NumGoroutine())

	go func() {
		fmt.Println("this is from go routine 1")
		wg.Done()

	}()

	fmt.Println(runtime.NumGoroutine())

	go func() {
		fmt.Println("this is from go routine 2")
		wg.Done()
	}()

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()

	fmt.Println(runtime.NumGoroutine())

}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//var wg sync.WaitGroup

func say(s string, wg *sync.WaitGroup) {
	fmt.Println("yeah, yeah", runtime.NumGoroutine())
	for i := 0; i < 5; i++ {
		//fmt.Println(runtime.NumGoroutine())
		//runtime.Gosched()
		//fmt.Println(s)
		fmt.Println(i, s)
	}
	wg.Done()

}

func main() {
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(1)
	go say("world", &wg)
	wg.Add(1)
	say("hello", &wg)
	wg.Wait()

}

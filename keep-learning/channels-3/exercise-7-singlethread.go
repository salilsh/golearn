package main

import (
	"fmt"
)

func main() {

	//ch := make(chan int)

	arr := []int{}

	for k := 1; k < 11; k++ {
		arr = append(arr, writeToChannel((k*10)+1)...)
		fmt.Println(arr)

	}

	//defer close(ch)

	//fmt.Println(arr)

	readValues(arr)

	fmt.Println("Exiting....")
}

func writeToChannel(end int) []int {
	a := []int{}
	for i := end - 10; i < end; i++ {
		a = append(a, i)
		//fmt.Println(a)
		//fmt.Println("run number: ", i)
		//time.Sleep(1 * time.Second)
	}

	return a

}

func readValues(arx []int) {
	for _, v := range arx {
		fmt.Println(v)
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
} */

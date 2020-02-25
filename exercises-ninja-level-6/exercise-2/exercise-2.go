package main

import (
	"fmt"
)

func tryOutSumUnfurl(y ...int) {
	sum := 0
	for _, value := range y {
		sum += value
	}

	fmt.Println(sum)
}

func tryOutSumDirectSlice(y []int) {
	sum := 0
	for _, value := range y {
		sum += value
	}

	fmt.Println(sum)

}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	tryOutSumUnfurl(x...)
	tryOutSumDirectSlice(x)

}

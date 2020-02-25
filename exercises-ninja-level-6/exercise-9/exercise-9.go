package main

import (
	"fmt"
)

func testCallback(fnc func(arr ...int) int, arr1 ...int) int {
	total := fnc(arr1...)
	return total

}


func main() {
	abcarr := []int{1, 2, 3, 4, 5, 6, 7}


	f := func(arr2 ...int) int {
		sum := 0
		for _, value := range arr2 {
			sum += value
		}
		return sum
	}

	finalTotal := testCallback(f, abcarr...)

	fmt.Println(finalTotal)
}

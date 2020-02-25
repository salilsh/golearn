package main

import (
	"fmt"
)

func trySendFuncReturn(x []int) func() int {

	total := 0
	fmt.Println(total, "this is at function init level. we have received", x)
	return func() int {
		for _, value := range x {
			total += value
		}

		return total
	}

}

func main() {
	fmt.Println("Hello, world")

	abc := []int{1, 2, 3, 4, 5, 6}

	g := trySendFuncReturn(abc)

	fmt.Println("Runnning trySendFuncReturn from within main")
	fmt.Println(g())
}

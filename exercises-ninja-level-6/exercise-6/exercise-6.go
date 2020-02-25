package main

import (
	"fmt"
)

func main() {
	var x int
	func() {
		x++
	}()

	fmt.Println("This my code's exec", x)
	fmt.Println()

	//Todd's code
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("Todd's code is done")
}

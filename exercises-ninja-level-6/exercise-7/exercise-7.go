package main

import (
	"fmt"
)

func main() {
	var g func()
	f := func() {
		for i := 0; i < 9; i++ {
			fmt.Println(i)

		}
	}

	f()

	g = f

	fmt.Println("now, this will run from g")
	g()

}

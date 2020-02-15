package main

import (
	"fmt"
)

func main() {
	for x := 1; x <= 10; x++ {
		//fmt.Println(x % 2)
		if (x % 2) == 0 {
			fmt.Println("the number is even", x)

		} else {
			fmt.Println("the number is odd", x)
		}

	}

}

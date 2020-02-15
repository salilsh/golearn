package main

import (
	"fmt"
)

func main() {
	x := 33
	for x < 123 {
		//y := fmt.Sprintf("%U", x)
		fmt.Printf("%U\n", x)
		x++
	}
}

package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func main() {
	x := 4
	y := 5
	z := 10

	defer fmt.Println(add(x, y))
	fmt.Println(subtract(x, y))
	fmt.Println(add(x, z))

	fmt.Println(add(z, y))

	fmt.Println(subtract(z, y))
}

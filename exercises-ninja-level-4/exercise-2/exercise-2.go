package main

import (
	"fmt"
)

func main() {

	slice := []int{2, 6, 6, 12, 24, 12, 15, 61, 19, 21}

	for r, val := range slice {
		fmt.Println(r, "->", val)

	}

	fmt.Printf("%T\n", slice)

}

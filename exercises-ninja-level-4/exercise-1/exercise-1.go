package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	for i := 0; i <= 4; i++ {
		arr[i] = i + 1
	}

	for r, val := range arr {
		fmt.Println(r, "->", val)

	}

	fmt.Printf("%T\n", arr)

}

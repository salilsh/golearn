package main

import (
	"fmt"
)

func main() {
	for charvalue := 65; charvalue <= 90; charvalue++ {
		fmt.Println(charvalue)
		fmt.Printf("\t\t%#U\n", charvalue)
		fmt.Printf("\t\t%#U\n", charvalue)
		fmt.Printf("\t\t%#U\n", charvalue)
	}

}

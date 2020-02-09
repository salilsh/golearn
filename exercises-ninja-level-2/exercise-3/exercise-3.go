package main

import (
	"fmt"
)

const (
	something          = 42
	typedSomething int = 42
)

func main() {
	fmt.Println(something)
	fmt.Println(typedSomething)
	fmt.Printf("%T\n", something)
	fmt.Printf("%T\n", typedSomething)
}

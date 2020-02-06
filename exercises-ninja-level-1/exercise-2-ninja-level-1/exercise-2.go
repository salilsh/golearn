package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println("value of int:", x)
	fmt.Println("value of string:", y)
	fmt.Println("value of bool:", z)
	outTypes()

}

func outTypes() {
	fmt.Printf("type of x is %T\n", x)
	fmt.Printf("type of y is %T\n", y)
	fmt.Printf("type of z is %T\n", z)
}

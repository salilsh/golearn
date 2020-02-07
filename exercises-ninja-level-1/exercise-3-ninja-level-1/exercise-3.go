package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println("value of int:", x)
	fmt.Println("value of string:", y)
	fmt.Println("value of bool:", z)
	s := fmt.Sprintf("%v\t%v\t%v",x,y,z) 
	fmt.Println(s)
	outTypes()

}

func outTypes() {
	fmt.Printf("type of x is %T\n", x)
	fmt.Printf("type of y is %T\n", y)
	fmt.Printf("type of z is %T\n", z)
}

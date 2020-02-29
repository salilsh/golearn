package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
}

func changeMe(p1 *person) {
	p1.firstname = "new guy"
	// (*p1).firstname = "new guy"
}

func main() {
	fmt.Println("Hello, world!")

	p1 := person{
		firstname: "salil",
		lastname:  "shankar",
	}

	fmt.Println(p1.firstname, p1.lastname)

	changeMe(&p1)

	fmt.Println(p1.firstname, p1.lastname)

}

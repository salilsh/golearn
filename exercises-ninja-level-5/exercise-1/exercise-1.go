package main

import (
	"fmt"
)

type person struct {
	firstname       string
	lastname        string
	iceCreamFlavors []string
}

func main() {
	p1 := person{
		firstname:       "salil",
		lastname:        "shankar",
		iceCreamFlavors: []string{"vanilla", "butterscotch", "tutti fruiti"},
	}

	p2 := person{
		firstname:       "vijay",
		lastname:        "chavan",
		iceCreamFlavors: []string{"chocolate", "chocolate", "chocolate"},
	}

	fmt.Println("entry:", p1.firstname, p1.lastname)
	for i, v := range p1.iceCreamFlavors {
		fmt.Printf("\tFavorite flavor number %v: %v\n", i, v)
	}

	fmt.Println()

	fmt.Println("entry:", p2.firstname, p2.lastname)
	for i, v := range p2.iceCreamFlavors {
		fmt.Printf("\tFavorite flavor number %v: %v\n", i, v)
	}

}

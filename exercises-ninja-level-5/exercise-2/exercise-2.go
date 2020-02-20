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
		iceCreamFlavors: []string{"pineapple", "vanilla", "strawberry"},
	}

	p2 := person{
		firstname:       "ayushi",
		lastname:        "sh",
		iceCreamFlavors: []string{"chocolate", "coffee", "butterscotch"},
	}

	superMap := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	for indi := range superMap {
		fmt.Println(superMap[indi].firstname)
		fmt.Println(superMap[indi].lastname)
		for i, value := range superMap[indi].iceCreamFlavors {
			fmt.Printf("\tFavorite icecream flavor number %v: %v\n", i, value)
		}
	}
}

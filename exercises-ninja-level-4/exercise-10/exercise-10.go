package main

import (
	"fmt"
)

func main() {
	peopleMap := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println("Printing the first map")
	fmt.Println()

	for r, inPtr := range peopleMap {
		fmt.Printf("%v\n", r)
		for r2, value := range inPtr {
			//fmt.Println()
			fmt.Printf("\tvalue at %v: %v\n", r2, value)
		}
		fmt.Println()

	}

	peopleMap["mallory_gareth"] = []string{`playing the boss`, `being smart`, `ralph fiennes`}

	fmt.Println("printing the updated map with new entry")
	fmt.Println()

	for r, inPtr := range peopleMap {
		fmt.Printf("%v\n", r)
		for r2, value := range inPtr {
			//fmt.Println()
			fmt.Printf("\tvalue at %v: %v\n", r2, value)
		}
		fmt.Println()

	}

	fmt.Println("printing the map with deleted entry")
	fmt.Println()

	delete(peopleMap, "no_dr")

	for key, inPtr := range peopleMap {
		fmt.Printf("%v\n", key)
		for r2, value := range inPtr {
			//fmt.Println()
			fmt.Printf("\tvalue at %v: %v\n", r2, value)
		}
		fmt.Println()

	}

}

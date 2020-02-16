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

	for r, inPtr := range peopleMap {
		fmt.Printf("%v\n", r)
		for r2, value := range inPtr {
			//fmt.Println()
			fmt.Printf("\tvalue at %v: %v\n", r2, value)
		}
		fmt.Println()

	}

}

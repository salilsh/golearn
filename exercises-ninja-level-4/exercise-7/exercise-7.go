package main

import (
	"fmt"
)

func main() {

	slice1 := []string{"James", "Bond", "Shaken, not stirred"}
	slice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	sony := [][]string{slice1, slice2}

	fmt.Println(sony)

	/*for r, value := range sony {
		fmt.Println(r, "->", value)
	}

	for r, value := range sony[i] {
		fmt.Println(r, "->", value)
	}

	for r, value := range sony[i+1] {
		fmt.Println(r, "->", value)
	}*/

	for r, slice := range sony {
		fmt.Println("record: ", r)
		for innerR, value := range slice {
			fmt.Printf("\t ")
			fmt.Println(innerR, "->", value)
		}
	}

	//this was for testing
	/*for r, value := range sony[i] {
		fmt.Printf("%v  %v  ", r, value)
		for r2, value2 := range sony[i+1] {
			fmt.Printf("%v  %v  ", r2, value2)
		}
		fmt.Println()
	}*/

}

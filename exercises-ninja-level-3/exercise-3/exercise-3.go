package main

import (
	"fmt"
)

func main() {
	yearOfBirth := 1988
	counter := -1
	for yearOfBirth < 2021 {
		fmt.Println(yearOfBirth)
		counter++
		fmt.Println(counter)
		yearOfBirth++
	}
	fmt.Println()
	fmt.Printf("You have been alive for %d years.\n", counter)
}

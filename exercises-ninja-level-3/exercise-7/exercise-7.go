package main

import (
	"fmt"
)

func main() {
	j := 0
	typeOfJ := fmt.Sprintf("%T", j)
	if typeOfJ == "bool" {
		fmt.Println("boolean")
	} else if typeOfJ == "string" {
		fmt.Println("string")
	} else {
		fmt.Println("something we don't know yet.")
	}
}

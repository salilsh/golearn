package main

import (
	"fmt"
)

func main() {
	person := struct {
		firstname string
		lastname  string
		age       int
	}{
		firstname: "salil",
		lastname:  "shankar",
		age:       31,
	}

	fmt.Println("The person is", person.firstname, person.lastname, "and his age is", person.age)
}

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	buffer, error := json.Marshal(users)

	if error != nil {
		fmt.Println("something is wrong", error)

	}

	fmt.Println(users)
	fmt.Println()
	fmt.Println(string(buffer))

	// your code goes here
}

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type multipleusers []user

func (u multipleusers) Len() int {
	return len(u)
}

func (u multipleusers) Less(i, j int) bool {
	if u[i].Age < u[j].Age {
		return true
	} else if u[i].Age == u[j].Age {
		if u[i].Last < u[j].Last {
			return true
		}
	}

	return false

}

func (u multipleusers) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	u4 := user{
		First: "M",
		Last:  "HMmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3, u4}

	fmt.Println(users)

	// your code goes here
	sort.Sort(multipleusers(users))

	for _, user := range users {
		sort.Strings(user.Sayings)
		fmt.Println(user.First, user.Last)
		fmt.Println(user.Age)
		for _, saying := range user.Sayings {
			fmt.Printf("\t%v\n", saying)
		}
		fmt.Println()

	}

}

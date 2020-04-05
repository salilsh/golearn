package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		return []byte{}, fmt.Errorf("There was an error in toJSON. The error was: %v", err)
		/* Or, these solutions by Todd:
			return []byte{}, fmt.Errorf("There was an error marshalling %v in toJSON: %v", a, err)
			return []byte{}, errors.New(fmt.Sprintf("There was an error in toJSON %v", err))
			^ requires importing errors library. which, VScode automatically does when you save your code.
		*/
	}

	return bs, nil
}

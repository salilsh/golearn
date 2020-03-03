package main

import (
	"fmt"
	"sort"
)

type person struct {
	name string
	age  int
}

type sortByName []person

func (s sortByName) Len() int {
	return len(s)
}

func (s sortByName) Swap(index1, index2 int) {
	s[index1], s[index2] = s[index2], s[index1]
}

func (s sortByName) Less(i, j int) bool {
	return (s[i].name < s[j].name)
}

func main() {

	var people []person

	people = []person{{"Q", 60}, {"Moneypenny", 27}, {"M", 71}, {"James", 32}}

	fmt.Println(people)
	sort.Sort(sortByName(people))
	fmt.Println()
	fmt.Println(people)
}

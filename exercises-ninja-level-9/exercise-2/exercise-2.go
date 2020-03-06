package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, "says hi! from the method")
	//a := p.first + " " + p.last
	//return a
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
	//fmt.Println(a, "from the interface")
}

func main() {
	h1 := person{
		first: "salil",
		last:  "shankar",
	}

	saySomething(&h1)

	//doesn't work. remove pointer receiver from the speak() function to make this work.
	//saySomething(h1)

}

package main

import (
	"fmt"
	"dog2"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog2.Dogyears(10),
	}
	fmt.Println(fido)
	fmt.Println(dog2.YearsTwo(20))
}
package main

import (
	"fmt"
)

func foo() int {
	return 3
}

func bar() (int, string) {
	return 3, "three"
}

func main(){
	x:=foo()
	a,b := bar()

	fmt.Println(x,a,b)

}
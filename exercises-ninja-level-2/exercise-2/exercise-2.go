package main

import (
	"fmt"
)

func main() {
	a := (42 == 42)
	b := (42 < 42)
	c := (42 >= 41)
	d := (44 != 44)
	e := (45 > 44)
	f := (46 <= 44)
	fmt.Println(a, b, c, d, e, f)

}

package main

import (
	"fmt"
)

func main() {
	x := make([]string, 50, 500)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// put a value into each of the 50 positions in the length of the slice
	// we are putting values into the 50 positions that are the length of the slice
	for i := 0; i < 50; i++ {
		x[i] = fmt.Sprintf("Position %d", i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// append values to the slice grows the length of the slice
	// the underlying array "capacity" of 500 is the same
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

//The question was misleading, and this program was written based on the understanding of the question
//the code above is a copy of the solution todd provided. he himself wrote a similar program to yours based on the understanding of his question
/*package main

import (
	"fmt"
)

func main() {
	statesOfUSA := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}
	fmt.Println(len(statesOfUSA))
	fmt.Println(cap(statesOfUSA))
	i := 0
	for i < len(statesOfUSA) {
		fmt.Println(i, "->", statesOfUSA[i])
		i++
	}
} */

package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here's the error: %v", ce.info)
}

func main() {
	tr := customErr{
		info: "It's a whole new world",
	}

	foo(tr)
}

func foo(e error) {
	fmt.Println("incoming from foo: ", e)
}

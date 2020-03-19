package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS is:", runtime.GOOS, "and architecture is:", runtime.GOARCH)
}

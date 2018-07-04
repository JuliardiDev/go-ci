package main

import (
	"fmt"
)

func main() {
	b := add(10, 20)
	fmt.Println(b)
}

// simple addition math func
func add(a, b int) int {
	return a + b
}

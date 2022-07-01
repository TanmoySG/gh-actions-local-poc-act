package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	sum := add(5, 6)
	fmt.Printf("%v\n", sum)
}

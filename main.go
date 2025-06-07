package main

import "fmt"

func add(a, b int) int {
	c := a + b
	return c
}

func main() {
	fmt.Printf("sum of %d and %d is %d\n", 5, 6, add(5, 6))
}

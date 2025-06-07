package main

import (
	"fmt"
)

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is:", result)

	newNumbers := []int{6, 7, 8, 9, 10}

	result = sum(newNumbers...)
	fmt.Println("The sum of new numbers is:", result)

}

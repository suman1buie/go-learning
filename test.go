package main

import "fmt"

func sumAnyNumbers(numbers ...int) int {

	sum := 0

	for num := range numbers {
		sum += num
	}
	return sum

}

func main() {
	fmt.Println(sumAnyNumbers(1, 1, 2))
	nums := []int{8, 1, 2, 9}

	fmt.Println(sumAnyNumbers(nums...))
}

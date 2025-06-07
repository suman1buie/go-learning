package main

import (
	"fmt"
)

func main() {

	userAge := make(map[string]int)
	userEmail := map[string]string{"sumit": "suuman@mail.com"}

	userAge["sumit"] = 20
	userAge["suman"] = 22
	fmt.Println("sumit age is", userAge["sumit"])
	fmt.Println("suman age is", userAge["suman"])
	fmt.Println("userAge map is", userEmail)

}

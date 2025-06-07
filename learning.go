package main

import (
	"fmt"
	"slices"
)

func array() {
	var names [5]string

	names[0] = "John"
	names[1] = "Jane"
	names[2] = "Doe"
	names[3] = "Alice"
	names[4] = "Bob"
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	for i, name := range names {
		fmt.Println(i, name)
	}

	age := [2]int{20, 30}
	fmt.Println("Age:", age)

	age1 := [...]int{3: 9, 6: 9, 7: 7}
	fmt.Println("Age:", age1)

}

func sliceexample() {

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func main() {
	sliceexample()
	// array()
	// var a int = 20;

	// switch a{
	// case 10:
	// 	fmt.Println("a is 10")
	// case 20:
	// 	fmt.Println("a is 20")
	// default:
	// 	fmt.Println("a is neither 10 nor 20")
	// }

	// whatAmI := func(i interface{}){
	// 	switch t := i.(type) {
	// 	case int:
	// 		fmt.Println("I am an int")
	// 	case bool:
	// 		fmt.Println("I am a bool")
	// 	default:
	// 		fmt.Println("I am of a different type", t)
	// 	}
	// }

	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")

}

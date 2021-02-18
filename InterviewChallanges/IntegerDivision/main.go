// Write a function to perform integer division without using either the / or *
// operators. Find a fast way to do it.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(IntegerDivision(10, 2))
	fmt.Println(IntegerDivision(100, 2))
	fmt.Println(IntegerDivision(10, 10))
	fmt.Println(IntegerDivision(192, 16))
	fmt.Println(IntegerDivision(12488, 2))
	fmt.Println(IntegerDivision(193, 16))
	fmt.Println(IntegerDivision(216, 16))
}

func IntegerDivision(x int, y int) (int, int) {
	// X divided by Y = result
	result := 0

	for x > 0 {
		x = x - y
		result++
	}

	fmt.Println("Integer 1 = result, Integer 2 = what's left")
	return result, -x
}

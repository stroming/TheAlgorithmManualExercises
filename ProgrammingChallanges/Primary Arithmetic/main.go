/*
Title Description

Children are taught to add multi-digit numbers from right-to-left one digit at a time.
Many find the "carry" operation - in which a 1 is carried from one digit position to be added to the next - to be a significant challenge.
Your job is to count the number of carry operations for each of a set of addition problems so that educators may assess their difficulty.

Entry
Each line of input contains two unsigned integers less than 10 digits. The last line of input contains 0 0.

Export
For each line of input except the last you should compute and print the number of carry operations
that would result from adding the two numbers.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// test cases
	fmt.Println(PrimaryArithmetic(12, 9))     // 1
	fmt.Println(PrimaryArithmetic(99, 9))     // 2
	fmt.Println(PrimaryArithmetic(999, 99))   // 3
	fmt.Println(PrimaryArithmetic(9999, 999)) // 4

}

func PrimaryArithmetic(one int, two int) (int, string) {
	// strTwo will always be smaller than strOne
	strOne := IntToString(one)
	strTwo := IntToString(two)
	if one < two {
		strOne = IntToString(two)
		strTwo = IntToString(one)
	}

	nBig := len(strOne)
	nSmall := len(strTwo)

	toBeAdded := 0
	result := 0
	for i, j := nBig-1, nSmall-1; i >= 0; i, j = i-1, j-1 {
		numOne := StringToInt(strOne[i])
		numTwo := toBeAdded
		if j >= 0 {
			numTwo += StringToInt(strTwo[j])
		}

		if numOne+numTwo >= 10 {
			result++
			toBeAdded = 1
		} else {
			toBeAdded = 0
		}
	}
	return result, " carry operations"
}
func IntToString(num int) string {
	return strconv.Itoa(num)
}
func StringToInt(str byte) int {
	result, _ := strconv.Atoi(string(str))
	return result
}

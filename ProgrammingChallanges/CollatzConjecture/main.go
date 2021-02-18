//  The Collatz conjecture is a conjecture in mathematics that concerns a sequence defined as follows:
//  Start with any positive integer n. Then each term is obtained from the previous term as follows:
//  if the previous term is even, the next term is one half of the previous term. If the previous term is odd, the next term is 3 times the previous term plus 1.
//  The conjecture is that no matter what value of n, the sequence will always reach 1.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(CollatzConjecture(1543))
	fmt.Println(CollatzConjecture(190))
	fmt.Println(CollatzConjecture(8))
	fmt.Println(CollatzConjecture(7))
	fmt.Println(CollatzConjecture(21890))
	fmt.Println(CollatzConjecture(13))
	fmt.Println(CollatzConjecture(451))
}

func CollatzConjecture(prevInt int) int {
	if prevInt == 1 {
		return 1
	}

	if prevInt%2 == 0 {
		return CollatzConjecture(prevInt / 2)
	}

	return CollatzConjecture(prevInt*3 + 1)
}

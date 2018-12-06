package main

import (
	"fmt"
)

/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

func main() {
	var fact100 = factorial(100)
	fmt.Println("The sum of the digits in 100! is ", sumDigitsInNum(fact100))
}

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func sumDigitsInNum(n int64) int64 {
	var sum int64

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}

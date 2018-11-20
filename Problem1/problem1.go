package main

import "fmt"

//Goal
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		if isMultiple(i, 3) || isMultiple(i, 5) {
			sum += i
		}
	}
	fmt.Println("The sum of multiples of 3 or 5 below 1000 is ", sum)

}

func isMultiple(a, b int) bool {
	if a%b == 0 {
		return true
	}
	return false
}

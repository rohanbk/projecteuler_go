package main

import (
	"fmt"
	"math/big"
)

/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

func main() {
	var fact100 = factorial(big.NewInt(100))
	fmt.Println("The sum of the digits in 100! is ", sumDigitsInNum(fact100))
}

func factorial(n *big.Int) *big.Int {
	//if n<2
	if n.Cmp(big.NewInt(2)) == -1 {
		return big.NewInt(1)
	} else {
		return Mul(n, factorial(big.NewInt(0).Sub(n, big.NewInt(1))))
	}
}

func sumDigitsInNum(n *big.Int) *big.Int {
	sum := big.NewInt(0)

	//for n > 0 {
	for big.NewInt(0).Cmp(n) == -1 {

		sum.Add(sum, big.NewInt(0).Mod(n, big.NewInt(10))) // sum+=n%10
		n = big.NewInt(0).Div(n, big.NewInt(10))           // n = n/10
	}

	return sum
}

// Mul helper function to multiply two big int values
func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

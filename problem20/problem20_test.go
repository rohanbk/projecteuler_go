package main

import (
	"math/big"
	"testing"
)

//TestFactorial test recursive factorial function
func TestFactorial(t *testing.T) {
	actualResult := factorial(big.NewInt(10))
	expectedResult := big.NewInt(3628800)

	if actualResult.Cmp(expectedResult) != 0 {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	actualResult = factorial(big.NewInt(1))
	expectedResult = big.NewInt(1)

	if actualResult.Cmp(expectedResult) != 0 {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	actualResult = factorial(big.NewInt(5))
	expectedResult = big.NewInt(120)

	if actualResult.Cmp(expectedResult) != 0 {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}
}

func TestSum(t *testing.T) {
	actualResult := sumDigitsInNum(big.NewInt(3628800))
	expectedResult := big.NewInt(27)

	if actualResult.Cmp(expectedResult) != 0 {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}
}

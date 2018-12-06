package main

import "testing"

//TestFactorial test recursive factorial function
func TestFactorial(t *testing.T) {
	actualResult := factorial(10)
	expectedResult := 3628800

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	actualResult = factorial(1)
	expectedResult = 1

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	actualResult = factorial(5)
	expectedResult = 120

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}
}

func TestSum(t *testing.T) {
	actualResult := sumDigitsInNum(3628800)
	expectedResult := 27
	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}
}

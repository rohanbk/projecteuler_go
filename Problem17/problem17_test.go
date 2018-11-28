package main

import "testing"

func TestProcessNumbersLessTen(t *testing.T) {
	actualResult := processNumbersLessTen(5)
	expectedResult := 4 //the word 'five'(5) is 4 characters long

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	actualResult = processNumbersLessTen(8)
	expectedResult = 5 //the word 'eight'(8) is 5 characters long

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}
}

func TestProcessNumbersInTeens(t *testing.T) {

	//the word 'eleven'(10) is 3 characters long
	actualResult := processNumbersInTeens(10)
	expectedResult := 3

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'twelve'(12) is 6 characters long
	actualResult = processNumbersInTeens(12)
	expectedResult = 6

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'sixteen'(16) is 7 characters long
	actualResult = processNumbersInTeens(16)
	expectedResult = 7

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'nineteen'(19) is 8 characters long
	actualResult = processNumbersInTeens(19)
	expectedResult = 8

	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}
}

func TestProcessNumbersLargerThanTwenty(t *testing.T) {

	//the word 'twenty one'(21) is 9 characters long
	actualResult := processNumbersLargerThanTwenty(21)
	expectedResult := 9
	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'sixty nine'(69) is 9 characters long
	actualResult = processNumbersLargerThanTwenty(69)
	expectedResult = 9
	if actualResult != expectedResult {
		t.Fatalf("Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'one hundred'(100) is 10 characters long
	actualResult = processNumbersLargerThanTwenty(100)
	expectedResult = 10
	if actualResult != expectedResult {
		t.Fatalf("100: Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'onehundredandone'(101) is 16 characters long
	actualResult = processNumbersLargerThanTwenty(101)
	expectedResult = 16
	if actualResult != expectedResult {
		t.Fatalf("101: Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'threehundredandfortytwo'(342) is 23 characters long
	actualResult = processNumbersLargerThanTwenty(342)
	expectedResult = 23
	if actualResult != expectedResult {
		t.Fatalf("342: Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'ninehundredandninetynine'(999) is 24 characters long
	actualResult = processNumbersLargerThanTwenty(999)
	expectedResult = 24
	if actualResult != expectedResult {
		t.Fatalf("999: Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'onethousand'(1000) is 11 characters long
	actualResult = processNumbersLargerThanTwenty(1000)
	expectedResult = 11
	if actualResult != expectedResult {
		t.Fatalf("1000: Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'onehundredfifteen'(115) is 20 characters long
	actualResult = processNumbersLargerThanTwenty(115)
	expectedResult = 20
	if actualResult != expectedResult {
		t.Fatalf("115: Expected %d but got %d", expectedResult, actualResult)
	}

	//the word 'onehundredten'(110) is 13 characters long
	actualResult = processNumbersLargerThanTwenty(110)
	expectedResult = 13
	if actualResult != expectedResult {
		t.Fatalf("110: Expected %d but got %d", expectedResult, actualResult)
	}
}

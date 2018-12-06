package problem17

import (
	"fmt"
)

var units = [10]string{"", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var teens = [10]string{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = [10]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

//HUNDRED string for the word hundred used for numbers > 100
const HUNDRED = "hundred"

//AND string for 'and' that acts as a delimiter for numbers > 100
const AND = "and"

var lengthOfNumberStrings map[int]int

func main() {
	processNumbers(1000)
}

func processNumbers(upperbound int) {
	lengthOfNumberStrings = make(map[int]int)
	sum := 0
	for i := 1; i <= upperbound; i++ {

		// initialize maps with lengths of corresponding strings
		if i < 10 {
			sum += processNumbersLessTen(i)
		} else if i >= 10 && i < 20 {
			sum += processNumbersInTeens(i)
		} else if i >= 20 {
			sum += processNumbersLargerThanTwenty(i)
		}
	}
	fmt.Println("The total sum is ", sum)
}

func processNumbersLessTen(n int) int {
	return len(units[n])
}

func processNumbersInTeens(n int) int {
	return len(teens[n%10])
}

func processNumbersLargerThanTwenty(num int) int {
	sum := 0
	i := 0
	n := num
	if n > 100 && n%100 != 0 {
		sum += len(AND)
	}

	for n > 0 {
		switch i {
		case 0:
			if num%100 < 10 || num%100 > 20 {
				x := units[n%10]
				sum += len(x)
			}

		case 1:
			if num%100 > 10 && num%100 < 20 {
				x := teens[num%10]
				sum += len(x)
			} else {
				x := tens[n%10]
				sum += len(x)
			}
		case 2:
			x := units[n%10]
			sum += len(x)
			if num < 1000 {
				sum += len(HUNDRED)
			}
		case 3:
			x := units[n%10]
			sum += len(x) + len("thousand")
		}
		i++
		n = n / 10
	}

	return sum
}

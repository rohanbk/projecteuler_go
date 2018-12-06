package problem10

import "fmt"

func main() {
	sieveOfEratothenes(2000000)
}

func sieveOfEratothenes(n int) {
	prime := make([]bool, n)
	for i, _ := range prime {
		prime[i] = true
	}

	p := 2

	for p*p <= n {
		// If prime[i] is not changed, then it is a prime
		if prime[p] {
			//Update all multiples of p
			for j := p * 2; j < n; j += p {
				if j%p == 0 {
					prime[j] = false
				}
			}
		}
		p++
	}
	//Iterate over prime slice and if the value is true, then print the index
	sum := 0
	for i := 2; i < len(prime); i++ {
		if prime[i] {
			sum += i
		}
	}
	fmt.Printf("The sume of prime numbers less than %d is %d\n", n, sum)
}

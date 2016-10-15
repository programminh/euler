package main

import "fmt"

func main() {
	var x, n int64

	n = 600851475143

	for {
		if isPrime(n) {
			fmt.Println("Largest Prime Factor", n)
			return
		}

		x = smallestPrimeFactor(n)
		n = n / x
	}
}

func smallestPrimeFactor(n int64) int64 {
	var i int64

	if isPrime(n) {
		return n
	}

	for i = 2; i < n/2; i++ {
		if n%i == 0 {
			return i
		}
	}

	return 0
}

func isPrime(n int64) bool {
	var i int64

	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i = 5

	for {
		if i*i > n {
			break
		}

		if n%i == 0 || n%(i+2) == 0 {
			return false
		}

		i += 6
	}

	return true
}

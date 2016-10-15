package main

import (
	"log"
	"strconv"
)

func main() {
	var x, m, a, b int

	big := 0
	n := 999

	for i := n; i >= 100; i-- {
		for m = 999; m >= 100; m-- {
			x = i * m
			if isPalindromicNumber(x) && x >= big {
				a = i
				b = m
				big = x
			}
		}
	}

	log.Println("Largest", big, a, b)
}

func isPalindromicNumber(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

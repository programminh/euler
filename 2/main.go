package main

import "fmt"

func main() {
	x := 0
	sum := 0
	i := 0

	for {
		x = fib(i)

		if x >= 4000000 {
			break
		}

		if x%2 == 0 {
			sum = sum + x
		}

		i++
	}

	fmt.Println("Sum is", sum)

}

func fib(n int) int {
	prev := 0
	current := 1

	for i := 0; i < n; i++ {
		current = current + prev
		prev = current - prev
	}

	return current
}

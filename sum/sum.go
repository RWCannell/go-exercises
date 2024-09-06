package main

import "fmt"

// sum of first n integers using iteration
func iterativeSum(n int) int {
	var sum int = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// sum of first n integers using recursion
func recursiveSum(n int) int {
	if n == 1 {
		return 1
	}
	return n + recursiveSum(n-1)
}

func main() {
	var n int = 100

	fmt.Printf("Iterative Sum of first %d integers: %d\n", n, iterativeSum(n))
	fmt.Printf("Recursive Sum of first %d integers: %d\n", n, recursiveSum(n))
}

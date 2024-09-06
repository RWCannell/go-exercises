package main

import "fmt"

// using head recursion
func headRecursionFactorial(n int) int {
	if n == 1 {
		return 1
	}
	var temp int = headRecursionFactorial(n - 1)
	var result int = n * temp
	return result
}

// using tail recursion
func tailRecursionFactorial(n int, acc int) int {
	if n == 0 {
		return acc
	}
	return tailRecursionFactorial(n-1, n*acc)
}

func main() {
	var n int = 5
	var accumulator int = 1

	var headRecursionFactorialValue = headRecursionFactorial(n)
	var tailRecursionFactorialValue = tailRecursionFactorial(n, accumulator)
	fmt.Printf("Head Recursion: The factorial of %d is %d\n", n, headRecursionFactorialValue)
	fmt.Printf("Tail Recursion: The factorial of %d is %d\n", n, tailRecursionFactorialValue)
}

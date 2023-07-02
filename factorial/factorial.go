package main

import "fmt"

// using head recursion
func head_recursion_factorial(n int) int {
	if n == 1 {
		return 1
	}
	var temp int = head_recursion_factorial(n - 1)
	var result int = n * temp
	return result
}

// using tail recursion
func tail_recursion_factorial(n int, acc int) int {
	if n == 0 {
		return acc
	}
	return tail_recursion_factorial(n-1, n*acc)
}

func main() {
	var n int = 7
	var accumulator int = 1

	var head_recursion_factorial_value = head_recursion_factorial(n)
	var tail_recursion_factorial_value = tail_recursion_factorial(n, accumulator)
	fmt.Printf("Head Recursion: The factorial of %d is %d\n", n, head_recursion_factorial_value)
	fmt.Printf("Tail Recursion: The factorial of %d is %d\n", n, tail_recursion_factorial_value)
}

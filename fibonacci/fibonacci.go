package main

import "fmt"

// function for calculating the nth Fibonacci number
func calculateNthFibonacciNumber(n int) int {
	if n == 0 || n==1 {
		return 1
	}
    return calculateNthFibonacciNumber(n-2) + calculateNthFibonacciNumber(n-1)
}

// return a list of Fibonacci numbers
func listFibonacciNumbers(n int) []int {
	var fibonacciNumber int = 0
	var fibonacciNumbers []int
    for i := 1; i <= n; i++ {
        fibonacciNumber = calculateNthFibonacciNumber(i)
		fibonacciNumbers = append(fibonacciNumbers, fibonacciNumber)
	}
    return fibonacciNumbers
}


func main() {
	var n int = 20
	var nthFibonacciNumber = calculateNthFibonacciNumber(n)
	var fibonacciNumbers = listFibonacciNumbers(n)
	fmt.Printf("nth Fibonacci number for n=%d is %d\n", n, nthFibonacciNumber)
	fmt.Printf("The first %d Fibonacci numbers are %d\n", n, fibonacciNumbers)
}
	
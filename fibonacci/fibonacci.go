package main

import "fmt"

// function for calculating the nth Fibonacci number
func calculate_nth_fibonacci_number(n int) int {
	if n == 0 || n==1 {
		return 1
	}
    return calculate_nth_fibonacci_number(n-2) + calculate_nth_fibonacci_number(n-1)
}

// return a list of Fibonacci numbers
func list_fibonacci_numbers(n int) []int {
	var fibonacci_number int = 0
	var fibonacci_numbers []int
    for i := 1; i <= n; i++ {
        fibonacci_number = calculate_nth_fibonacci_number(i)
		fibonacci_numbers = append(fibonacci_numbers, fibonacci_number)
	}
    return fibonacci_numbers
}


func main() {
	var n int = 10
	var nth_fibonacci_number = calculate_nth_fibonacci_number(n)
	var fibonacci_numbers = list_fibonacci_numbers(n)
	fmt.Printf("nth Fibonacci number for n=%d is %d\n", n, nth_fibonacci_number)
	fmt.Printf("The first %d Fibonacci numbers are %d\n", n, fibonacci_numbers)
}
	
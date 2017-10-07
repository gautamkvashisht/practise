package main

// Take N as input and print the number at Nth positon in the Fibonacci series

import "fmt"

func main() {
	var N int
	first := 0
	second := 1
	count := 1
	fmt.Print("Enter the position at which you want to know the value of Fibonacci series (if 0 is 0th Fibonacci number and 1 is 1st Fibonacci number): ")
	fmt.Scanln(&N)
	if N == 0 {
		count = 0
	} else {
		for i := 1; i < N; i++ {
			count = first + second
			first = second
			second = count
		}
	}
	fmt.Println("The number at", N, "th position is", count)
}

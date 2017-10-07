package main

import "fmt"

// Print all Fibonacci numbers less than or equal to a certain number

func main() {
	var N int
	first := 0
	second := 1
	fmt.Print("Enter the number upto which you want to find Fibonacci series: ")
	fmt.Scan(&N)
	fmt.Println(0)
	for count := 1; count <= N; {
		fmt.Println(second)
		count = first + second
		first = second
		second = count
	}
}

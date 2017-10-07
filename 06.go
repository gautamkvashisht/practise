package main

import "fmt"

// Take N as input. Print the sum of its odd digits and sum of its even digits

func main() {
	var N int
	evensum := 0
	oddsum := 0
	i := 0
	fmt.Print("Enter the number: ")
	fmt.Scan(&N)
	for ; N != 0; N = N / 10 {
		remainder := N % 10
		if (remainder % 2) == 0 {
			evensum += remainder
		} else {
			oddsum += remainder
		}
		i++
	}
	fmt.Println("Sum of even digits= ", evensum)
	fmt.Println("Sum of odd digits= ", oddsum)
}

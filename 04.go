package main

import "fmt"

// Take N as input and print the prime numbers from 2 to N

func main() {
	var N int
	k := 0
	fmt.Print("Enter the number upto which you want to know the prime numbers: ")
	fmt.Scan(&N)
	if N == 2 {
		fmt.Println(N)
	} else {
		fmt.Println(2)
		for i := 3; i <= N; i++ {
			flag := 0
			for j := 2; j <= (i+1)/2; j++ {
				if (i % j) == 0 {
					flag++
				}
			}
			if flag == 0 {
				k++
				fmt.Println(i)
			}
		}
	}
}

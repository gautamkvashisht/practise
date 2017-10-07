package main

// Take a number N as input and check whether it is prime or not

import "fmt"

func main() {
	var N int
	flag := 0
	fmt.Print("Enter the positive integer (other than 1): ")
	fmt.Scanln(&N)
	if N == 2 {
		fmt.Println(N, "is Prime")
		flag = 1
	} else {
		for i := 2; i <= (N / 2); i++ {
			if (N % i) == 0 {
				flag++
			}
		}
		if flag > 0 {
			fmt.Println(N, "is Not Prime")
		} else {
			fmt.Println(N, "is Prime")
		}
	}
}

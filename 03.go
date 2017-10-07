package main

import "fmt"

// Take a string as an input and print its reverse

func main() {
	var input string
	fmt.Print("Enter the input: ")
	fmt.Scanln(&input)
	arr := []byte(input)
	x := len(input)
	for i, j := 0, x-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	input = string(arr[:])
	fmt.Println(input)
}

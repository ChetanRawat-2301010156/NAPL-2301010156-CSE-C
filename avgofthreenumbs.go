package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Print("Enter three numbers: ")
	fmt.Scan(&a, &b, &c)

	average := (a + b + c) / 3

	fmt.Println("Average:", average)

	if average >= 50 {
		fmt.Println("Result: Pass")
	} else {
		fmt.Println("Result: Fail")
	}
}

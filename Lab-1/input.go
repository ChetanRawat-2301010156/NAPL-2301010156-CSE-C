package main

import "fmt"

func main() {
	var marks int

	fmt.Print("Enter your marks (0-100): ")
	fmt.Scan(&marks)

	if marks < 0 || marks > 100 {
		fmt.Println("Invalid marks!")
	} else {
		fmt.Println("Valid marks:", marks)
	}
}

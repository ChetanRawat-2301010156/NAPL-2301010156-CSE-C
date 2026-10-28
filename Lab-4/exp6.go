package main

import (
	"fmt"
	"sync"
)

// Calculates the square of a number.
func calculateSquare(n int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Square goroutine started")

	result := n * n
	ch <- fmt.Sprintf("Square of %d = %d", n, result)

	fmt.Println("Square goroutine completed")
}

// Calculates the cube of a number.
func calculateCube(n int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Cube goroutine started")

	result := n * n * n
	ch <- fmt.Sprintf("Cube of %d = %d", n, result)

	fmt.Println("Cube goroutine completed")
}

// Calculates the first n Fibonacci numbers.
func calculateFibonacci(n int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Fibonacci goroutine started")

	a, b := 0, 1
	sequence := ""

	for i := 0; i < n; i++ {
		sequence += fmt.Sprintf("%d ", a)
		a, b = b, a+b
	}

	ch <- fmt.Sprintf("First %d Fibonacci numbers: %s", n, sequence)

	fmt.Println("Fibonacci goroutine completed")
}

func main() {
	results := make(chan string, 3)
	var wg sync.WaitGroup

	wg.Add(3)

	// Launch three independent goroutines.
	go calculateSquare(5, results, &wg)
	go calculateCube(3, results, &wg)
	go calculateFibonacci(7, results, &wg)

	// Collect and print all three results.
	for i := 0; i < 3; i++ {
		fmt.Println(<-results)
	}

	wg.Wait()
	close(results)
}

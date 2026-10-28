
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Scenario 1: Handle division by zero.
func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}

// Scenario 2: Handle negative square root input.
func safeSqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("cannot calculate square root of %v", n)
	}
	return math.Sqrt(n), nil
}

// Scenario 3: Handle invalid positive integer input.
func parsePositiveInt(input string) (int, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return 0, fmt.Errorf("input cannot be empty")
	}

	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid integer: %w", err)
	}

	if n <= 0 {
		return 0, fmt.Errorf("number must be positive")
	}

	return n, nil
}

func main() {
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	root, err := safeSqrt(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square root:", root)
	}

	number, err := parsePositiveInt("10")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed integer:", number)
	}

	// Demonstrate the three error scenarios.
	_, err = safeDivide(10, 0)
	fmt.Println("Division error:", err)

	_, err = safeSqrt(-4)
	fmt.Println("Square root error:", err)

	_, err = parsePositiveInt("abc")
	fmt.Println("Input error:", err)
}

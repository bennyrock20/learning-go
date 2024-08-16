package calculator

import "fmt"

// Define a Calculator struct
type Calculator struct{}

// Method to add two numbers
func (c Calculator) Add(a, b int) int {
	return a + b
}

// Method to subtract two numbers
func (c Calculator) Subtract(a, b int) int {
	return a - b
}

// Method to multiply two numbers
func (c Calculator) Multiply(a, b int) int {
	return a * b
}

// Method to divide two numbers
func (c Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

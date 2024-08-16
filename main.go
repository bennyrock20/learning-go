package main

import (
	"fmt"
	"learning-go/utils/calculator"
)

func main() {

	calc := calculator.Calculator{}

	a := 4
	b := 2

	sum := calc.Add(a, b)
	fmt.Println("Sum", sum)
	fmt.Printf("a: %d, b: %d, sum: %d\n", a, b, sum)

	sub := calc.Subtract(4, 2)
	fmt.Printf("a: %d, b: %d, sub: %d\n", a, b, sub)

	mult := calc.Multiply(a, b)
	fmt.Printf("a: %d, b: %d, mult: %d\n", a, b, mult)

	div, err := calc.Divide(a, b)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("a: %d, b: %d, div: %d\n", a, b, div)
	}

}

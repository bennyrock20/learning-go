package main

import (
	"fmt"
	"learning-go/utils/calculator"
	"learning-go/utils/person"
	"learning-go/utils/webserver"
)

func run_go() {
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

	// Another test

	p1 := person.Person{
		FirstName: "Wil",
		LastName:  "Tapia",
		Age:       35,
	}

	fmt.Println("Full Name:", p1.FullName())
	fmt.Println("Is Adult", p1.IsAdult())

	// Test "Constructor Function"

	p2 := person.NewPerson("Jesse", "Guncay", 34)
	// Test String methoda
	fmt.Print(p2)

	// print json
	fmt.Println(p2.ToJSON())
}

func run_server() {
	webserver.Up()

}
func main() {

	// run_go()

	// Testin web server

	run_server()

}

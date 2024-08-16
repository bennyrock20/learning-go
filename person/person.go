package person

import (
	"encoding/json"
	"fmt"
)

// Define a struct to represent a Person
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

// Constructor function returning a pointer to Person
func NewPerson(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// String method to convert the Person struct to a string.
func (p Person) String() string {
	return fmt.Sprintf("Person: %s %s, Age: %d", p.FirstName, p.LastName, p.Age)
}

// ToJSON converts the Person struct to a JSON string.
func (p *Person) ToJSON() (string, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", fmt.Errorf("error marshalling to JSON: %v", err)
	}
	return string(jsonData), nil
}

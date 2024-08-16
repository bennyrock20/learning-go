package person

// Define a struct to represent a Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

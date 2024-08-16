package person

// Define a struct to represent a Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
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

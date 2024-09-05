package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) { // A pointer must be received in order to mutate the value
	(*p).firstName = newFirstName
}

func main() {
	// var alex person // Definition without initialization
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex := person{"Alex", "Anderson"} // Bad code!
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email: "alex@mail.com",
			zip:   94000,
		},
	} //Better!

	// alexPtr := &alex // You can declare a pointer like this

	/* In Go arguments are pass by value, which means that a copy of them is created when passing them to a function.

	You can mutate a variable of type T in a function by letting Go automatically convert it to a pointer
	if the argument's type is of 'pointer to T'.

	Slices don't work in the same way though! Because they contain a pointer to the underlying array,
	modifying a copy of them in a function will result on the original array value being changed too!

	This is the main difference between value types and reference types:
	Value types: int, float, string, bool, struct
	Reference types: slices, maps, channels, pointers, functions
	*/
	alex.updateName("Jimmy")

	alex.print()
}

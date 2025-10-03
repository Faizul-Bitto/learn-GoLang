//! In this section we will be using 'struct' to create a 'receiver' function. So, it's recommended to first watch the 'struct' section to understand the concept of 'receiver' function

package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

//! we know to define a parameter in a function, we have to also mention the data type of the parameter. Here, 'Person' is the data type of the parameter. 'person' is the parameter and 'Person' is the type. Just like we define : (name string, age int, email string) as parameters
func printUserDetails(person Person) {
	fmt.Println(`Person Name :`, person.Name, `Person Age :`, person.Age, `Person Email :`, person.Email)

}

//! just like same, if we just change the structure of the function like this :
func (person Person) printDetails() {
	fmt.Println(`Person Name :`, person.Name, `Person Age :`, person.Age, `Person Email :`, person.Email)

} //! this is a receiver function. this structure is only possible with the custom data type made with 'struct' keyword.

func main() {
	var person1 Person

	person1 = Person{
		Name:  "John",
		Age:   20,
		Email: "john@example.com",
	}

	printUserDetails(person1)

	//! now for that receiver function, we can call it like this :
	person1.printDetails()

	person2 := Person{
		Name:  "Jane",
		Age:   21,
		Email: "jane@example.com",
	}

	printUserDetails(person2)
}

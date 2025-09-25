//! As there's many types of data. While declaring variables, we can set the data type of the variable. But, we can also make our own data type. For example, we can make a data type of a person. A person has a name, age, and email. And, we do it with 'struct'.

package main

import "fmt"

//! first write 'type' keyword, then write the name of the struct, then write the fields of the struct in curly braces.
type Person struct {
	Name  string //! member variable or property
	Age   int    //! member variable or property
	Email string //! member variable or property
} //! so, here, Person is a data type. and, it has 3 fields : Name, Age, and Email.

func main() {
	//! we can declare a variable of type Person
	//! 'var' keyword then the 'variable name' and then the data type, which is 'Person' in this case
	var person Person
	person = Person{
		Name:  "John",
		Age:   20,
		Email: "john@example.com",
	}
	fmt.Println(`Person Name :`, person.Name, `Person Age :`, person.Age, `Person Email :`, person.Email)

	//! we can also declare a variable of type Person using a short variable declaration -> instantiation of Person object
	person2 := Person{ //! person2 is the instance of Person object
		Name:  "Jane",
		Age:   21,
		Email: "jane@example.com",
	}
	fmt.Println(`Person Name :`, person2.Name, `Person Age :`, person2.Age, `Person Email :`, person2.Email)
}

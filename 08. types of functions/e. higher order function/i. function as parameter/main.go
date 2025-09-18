//! 'first order function' and 'higher order function' concepts came from 'functional paradigm'
//! there are two types of logics : 'first order logic' and 'higher order logic'

//! let's come to logics:

/*
	1. object
	2. property
	3. relation

	Let's think about a logic :
	Rule : All customers must pay their pizza bills.
	Object : Customer

	Tutul is a student.
	Object : Tutul
	Property : Student
	Relation : is

	Like this with these 3 properties, we can define any statements.

	Another Rule let's say :
	All students must wear their Identity Cards.

	These are first order logics:
	All customers must pay their pizza bills.
	All students must wear their Identity Cards.
	Tutul is a student.

	So, first order logics work with objects, properties, and relations.
*/

/*
	1. object
	2. property
	3. relation

	higher logic order works with rules

	Rule : Any rules that apply to all customers must also apply to all students.
	Higher Order Logics not only just work with object, property, relations but also work with rules.
*/

package main

import "fmt"

//! higher order functions :
/*
	To be a higher order function, a function must have at least one of these features:
	1. function as a parameter
	2. function as a return value
	3. both
*/

//! function as a parameter
func higherOrderFunction(a int, b int, f func(x int, y int) int) {
	fmt.Println(f(1, 2))
	fmt.Println(f(a, b))
}

func main() {
	higherOrderFunction(1, 2, calculateAdd)
}

func calculateAdd(a, b int) int {
	return a + b
}

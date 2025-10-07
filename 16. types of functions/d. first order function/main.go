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

//! first order functions :

/*
	first order function types:

	1. named/standard function
	2. anonymous function / IIFE

	higher order functions also known as 'first class functions' -> treated as 'first class citizens'
*/

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//! data types

	//! Numeric Data Types
	//! Integer Data Types
	a := 10
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println("--------------------------------")

	//! Floating-Point Data Types
	b := 10.5
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println("--------------------------------")

	//! Boolean Data Type
	d := true
	fmt.Println(d)
	fmt.Println(reflect.TypeOf(d))
	fmt.Println("--------------------------------")

	//! String Data Type
	e := "Hello World"
	fmt.Println(e)
	fmt.Println(reflect.TypeOf(e))
	fmt.Println("--------------------------------")
}

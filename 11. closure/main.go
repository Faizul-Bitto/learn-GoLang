package main

import "fmt"

const a = 10

var p int = 20

func outer() func() {
	money := 100
	age := 20

	fmt.Println("age =", age)

	show := func() {
		money = money + a + p
		fmt.Println("money =", money)
	}
	//! This 'show' function will be vanished after the outer function is executed. So, there will be no track or unable to use even we have returned this function. But, we can use this. How? Because there's a concept called closure, Heap and Garbage Collector. Go, compiler will understand that in the compilation time that, this outer() function's properties can be used in the future. So, it will be stored in the Heap memory. This is done by compiler's 'escape analysis' feature. And, we can use this function in the future. And this is the concept of closure. So, this show() function, 'money' will be stored in the Heap memory. Question can arise, why not 'a' and 'p'? Because 'a' is a constant variable and 'p' is a normal variable. So, 'a'' will be stored in the 'Code Segment'. And, 'p' will stored in the 'Data Segment' of the memory. They will not be cleaned up. And then after finishing work, if Garbage Collector understands that, the Heap will not used, then Garbage Collector will clean up the Heap memory.

	return show
}

func call() {
	increment1 := outer()
	increment1()
	increment1()
	increment1()

	increment2 := outer()
	increment2()
	increment2()
	increment2()
}

func main() {
	call()
}

func init() {
	fmt.Println("init function")
}

/*
	Now build with : go build -gcflags="-m" main.go
*/

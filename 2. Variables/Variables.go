package main

import "fmt"

func main() {
	// 1. var keyword
	var x int = 10
	var y string = "Sid-80"
	fmt.Println(x, y)

	//2. Declaring multiple variables
	var (
		firstName, Id string
		age           int
	)
	firstName = "Siddharth"
	Id = "Sid-80"
	age = 21
	fmt.Println(firstName, Id, age)

	//3. Short declaration
	Lang := "GO!"
	fmt.Println(Lang)

	// Default Values
	// 1. 0 for numeric values
	// 2. false for boolean values
	// 3. "" for string values
	// 4. <nil> for pointers

	// 4. constant variables : can't changed
	const PI = 3.14
	fmt.Println(PI)

	// Variable types
	// 1. Basic Types : int, float64, string, bool, etc.
	// 2. Composite types : array, struct, interface, slice, map, channel, etc.
}

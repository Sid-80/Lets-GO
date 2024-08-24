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
	age = 22
	fmt.Println(firstName, Id, age)

	//3. Short declaration
	Lang := "GO!"
	fmt.Println(Lang)
}

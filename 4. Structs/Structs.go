package main

import "fmt"

func main() {
	type Person struct {
		Name string
		age  int
	}

	var p1 Person = Person{"Sid", 21}

	fmt.Println(p1)
	fmt.Println("Name : ", p1.Name)
	fmt.Println("Age : ", p1.age)
}

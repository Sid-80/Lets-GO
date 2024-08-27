package main

import "fmt"

func main() {
	r := map[string]int{"Sid": 21}
	r["John"] = 18
	fmt.Println(r)
	fmt.Println("Age : ", r["Sid"])
	fmt.Println()

	// To delete
	delete(r, "John")
	fmt.Println("After Deleting : ", r)

	// A two-value assignment tests for the existence of a key:

	i, ok := r["John"]

	if !ok {
		fmt.Println("Deleted Successfully !")
	} else {
		fmt.Println("Age : ", i)
	}

	// Iterate
	fmt.Println()
	for key, value := range r {
		fmt.Println("Key:", key, "Value:", value)
	}

}

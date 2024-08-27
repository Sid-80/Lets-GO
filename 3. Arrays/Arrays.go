package main

import "fmt"

func main() {
	// Arrays :Static size
	var arr [5]int = [5]int{1, 2, 5}

	fmt.Println(arr)

	// Slices : Dynamic size
	arr2 := []int{1, 2, 3, 4}

	fmt.Println(arr2)

	// Creating an array whose size is determined
	// by the number of elements present in it
	// Using ellipsis
	myarray := [...]string{"Sid", "loves", "playing",
		"Basketball"}

	fmt.Println("Elements of the array: ", myarray)

	// Length of the array
	// is determine by
	// Using len() method
	fmt.Println("Length of the array is:", len(myarray))

	// Note : In Go language, an array is of value type not of reference type.
	// So when the array is assigned to a new variable, then the changes made in the new variable do not affect the original array.

}

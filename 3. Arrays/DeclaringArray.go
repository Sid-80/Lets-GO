package main

import "fmt"

func main() {
	// 1. Without init
	var arr [5]int

	// 2. Declaring with init
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}

	// 3. Shorthand
	arr3 := [5]int{1, 2, 3, 4, 5}

	// 4. specific element
	arr4 := [5]int{1: 20, 3: 40}

	// 5. ... Operator to infer length
	arr5 := [...]int{1, 2, 3, 4, 5}

	// // 6. Multi-Dimensional Array
	// var multiArr [2][3]int

	// //7. init multi
	// multiArr2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// 8. Arrays of pointers
	var arr6 [3]*int //This declares an array of pointers to integers var multiArr [2][3]int
	n1, n2, n3 := 11, 12, 13
	arr6[0] = &n1
	arr6[1] = &n2
	arr6[2] = &n3

	for i := 0; i < len(arr6); i++ {
		fmt.Println(*arr6[i])
	}
	for _, value := range arr6 {
		fmt.Println(*value)
	}

	// 9. Array of Struct
	type Person struct {
		Name string
		Age  int
	}

	// var people [2]Person

	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

}

package main

import "fmt"

func main() {
	var v func(int) int
	v = func(x int) int { return x * x }

	result := v(5)

	fmt.Println(result)
}

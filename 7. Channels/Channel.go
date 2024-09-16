// Goroutine Communication :
// Channels can be thought of as pipes using which Goroutines communicate.
// Similar to how water flows from one end to another in a pipe, data can be sent from one end and received from the other end using channels.
// Used for async or sync communication
package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
}

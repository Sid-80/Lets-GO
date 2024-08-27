package main

func main() {
	var s int = 10
	var t *int = &s

	println("s : ", s)
	println("t : ", t)
	println("value at t : ", *t)

	*t = 20
	println("s : ", s)
	println("t : ", t)
	println("value at t : ", *t)
}

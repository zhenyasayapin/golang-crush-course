package main

import "fmt"

func main() {
	x := 10
	pointer := &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", pointer)

	*pointer = 20
	fmt.Println("New value of x:", x)

	changeValue(pointer)
	fmt.Println("Value of x after changeValue:", x)
}

func changeValue(val *int) {
	*val = 30
}

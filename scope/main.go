package main

import (
	"fmt"
	"myapp/packageone"
)

func main() {
	var one = "One"

	fmt.Println(one)

	myFunction()

	newString := packageone.PublicVar
	fmt.Println("From package one", newString)

	packageone.PublicFunction()
}

func myFunction() {
	var one = "One from myFunction"
	fmt.Println(one)
}

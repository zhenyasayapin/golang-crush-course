package main

import "myapp/packageone"

var myVar = "It is main package variable"

func main() {
	var blockVar = "I am block scoped variable"

	packageone.PrintMe(myVar, blockVar)
}

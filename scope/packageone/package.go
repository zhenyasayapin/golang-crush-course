package packageone

import "fmt"

var PackageVar = "I am package one variable"

func PrintMe(myVar, blockBar string) {
	fmt.Println(PackageVar, myVar, blockBar)
}

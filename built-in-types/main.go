package main

import "log"

var myInt int
var myUint uint

var myFloat32 float32
var myFloat64 float64

func main() {
	myInt = 10
	myUint = 20

	myFloat32 = 30.5
	myFloat64 = 40.7

	log.Println(myInt, myUint, myFloat32, myFloat64)

	myString := "Hello, Go!"

	log.Println(myString)

	var myBool = true
	myBool = false

	log.Println(myBool)
}

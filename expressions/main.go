package main

import "fmt"

func main() {
	age := 30
	name := "Alice"
	rightHanded := true

	rightHanded = false

	fmt.Printf("%s is %d years old. Right handed: %t\n", name, age, rightHanded)

	ageInThenYears := age + 10

	fmt.Printf("In then years %s will be %d years old", name, ageInThenYears)
	fmt.Println()

	isATeenager := age >= 13

	fmt.Printf("Is %s a teenager? %t\n", name, isATeenager)
}

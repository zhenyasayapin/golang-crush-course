package main

import "fmt"

type Cat struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type Dog struct {
	Name         string
	Sound        string
	NumberOfLegs int
}

type AnimalInterface interface {
	Says() string
	HowManyLegs() int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (c *Cat) Says() string {
	return c.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

func (c *Cat) HowManyLegs() int {
	return c.NumberOfLegs
}

func main() {
	dog := Dog{
		Name:         "Buddy",
		Sound:        "Woof",
		NumberOfLegs: 4,
	}

	cat := Cat{
		Name:         "Whiskers",
		Sound:        "Meow",
		NumberOfLegs: 4,
	}

	Riddle(&dog)
	Riddle(&cat)
}

func Riddle(animal AnimalInterface) {
	fmt.Printf("The animal says: %s\n", animal.Says())
	fmt.Printf("It has %d legs\n", animal.HowManyLegs())
	fmt.Println("What animal is it?")
}

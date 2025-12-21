package main

import "fmt"

func main() {
	apples := 10
	oranges := 15

	fmt.Println(apples == oranges)
	fmt.Println(apples != oranges)

	fmt.Printf("Are apples > oranges: %t", apples > oranges)
	fmt.Println()

	fmt.Printf("Are apples < oranges: %t", apples < oranges)
	fmt.Println()

	fmt.Printf("Are apples >= oranges: %t", apples >= oranges)
	fmt.Println()

	fmt.Printf("Are apples <= oranges: %t", apples <= oranges)
	fmt.Println()
}

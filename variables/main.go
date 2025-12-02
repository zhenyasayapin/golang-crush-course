package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER..."

func main() {
	var firstNumber = 10
	var secondNumber = 2
	var subtraction = 7
	// var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number Game!")
	fmt.Println("______________________")
	fmt.Println("")
	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by ", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Nom multiply your result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now divide your result by your original number", prompt)
	reader.ReadString('\n')

	fmt.Println("Finally subtract", subtraction, prompt)
	reader.ReadString('\n')
}

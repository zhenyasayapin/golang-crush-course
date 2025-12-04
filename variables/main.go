package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER..."

func main() {
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2
	var subtraction = rand.Intn(8) + 2
	var answer int

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the Number Game!")
	fmt.Println("______________________")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Nom multiply your result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now divide your result by your original number", prompt)
	reader.ReadString('\n')

	fmt.Println("Finally subtract", subtraction, prompt)
	reader.ReadString('\n')

	answer = (firstNumber * secondNumber) - subtraction

	fmt.Println("The answer is:", answer)
}

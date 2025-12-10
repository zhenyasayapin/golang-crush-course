package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	Name           string
	Age            int
	FavoriteNumber float64
	OwnsADog       bool
}

func main() {
	reader = bufio.NewReader(os.Stdin)

	var user User

	user.Name = readString("Enter your name:")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your favorite number?")

	fmt.Printf("Your name is %s. You are %d years old. Your favorite number is %.2f.\n",
		user.Name,
		user.Age,
		user.FavoriteNumber)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "" {
			fmt.Println("Enter a valid value")
		} else {
			return userInput
		}
	}
}

func readInt(s string) int {
	fmt.Println(s)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	num, err := strconv.Atoi(userInput)

	if err != nil {
		fmt.Println("Please enter an integer.")
		return readInt(s)
	}

	return num
}

func readFloat(s string) float64 {
	fmt.Println(s)
	prompt()
	userInput, _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\r\n", "", -1)
	userInput = strings.Replace(userInput, "\n", "", -1)

	num, err := strconv.ParseFloat(userInput, 64)

	if err != nil {
		fmt.Println("Please enter a floating point number.")
		return readFloat(s)
	}

	return num
}

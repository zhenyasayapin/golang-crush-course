package main

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
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
	user.OwnsADog = readBool("Do you own a dog? (y/n)")

	fmt.Printf(
		"Your name is %s. You are %d years old. Your favorite number is %.2f. Owns a dog: %t\n",
		user.Name,
		user.Age,
		user.FavoriteNumber,
		user.OwnsADog,
	)
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

func readBool(s string) bool {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		fmt.Println(s)
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if strings.ToLower(string(char)) != "y" && strings.ToLower(string(char)) != "n" {
			fmt.Println("Please, use only y/n")
		} else if char == 'y' || char == 'Y' {
			return true
		} else if char == 'n' || char == 'N' {
			return false
		}
	}
}

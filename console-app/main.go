package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func main() {
	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	coffees := make(map[int]string)
	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("_____")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Espresso")
	fmt.Println("q - Quite the program")

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal(err)
		}

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))

		fmt.Printf("You chose %s\n", coffees[i])
	}

	fmt.Println("Programm exiting...")
}

package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func main() {
	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press anykey or press 'q' to quit")
	_ = keyboard.Open()

	defer keyboard.Close()

	key, _, _ := keyboard.GetSingleKey()

	for key != 'q' {
		key, _, _ = keyboard.GetSingleKey()
		keyPressChan <- key
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed:", string(key))
	}
}

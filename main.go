package main

import "fmt"

func main() {
	whatToSay := "Bonjour, tout le monde!"

	say(whatToSay)
}

func say(whatToSay string) {
	fmt.Println(whatToSay)
}

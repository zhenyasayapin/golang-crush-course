package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	i := 1000

	for i > 100 {
		i = rand.Intn(1000) + 1
		fmt.Println(i, "is the current value")
		if (i > 100) {
			fmt.Println("Still more than 100")
		}
	}

	fmt.Println("Got it! The value is", i)
}
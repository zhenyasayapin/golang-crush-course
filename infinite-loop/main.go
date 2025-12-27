package main

import (
	"bufio"
	"fmt"
	"mylogger/mylogger"
	"os"
	"time"
)

func main() {
	ch := make(chan string)

	go mylogger.ListenForLog(ch)

	fmt.Println("Enter something")

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 5; i++ {
		fmt.Print("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	for {
		var n int
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Input no of seconds and reminder message..")
		fmt.Scanln(&n)
		message, _ := reader.ReadString('\n')
		fmt.Println("Setting up reminder...")
		timer := time.NewTimer(time.Duration(n) * time.Second)
		<-timer.C
		fmt.Println("Hey!! ", message)
	}
}

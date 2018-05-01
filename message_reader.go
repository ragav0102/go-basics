package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ExitText = "exit\n"
)

func HandleMessage(message string) {
	fmt.Println("Message received: ", message)
}

func receiver(connection chan string) {
	fmt.Println("Setting up receiver...")
	for {
		fmt.Println("Waiting for message")
		HandleMessage(<-connection)
	}
	fmt.Println("Done in receiver")
}

func main() {
	channel := make(chan string)
	reader := bufio.NewReader(os.Stdin)

	go receiver(channel)

	fmt.Println("Ready to accept messages!!")

	for {
		input, _ := reader.ReadString('\n')
		if input == ExitText {
			break
		} else {
			channel <- input
		}
	}

	fmt.Println("Done handling messages")
}

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

func ReadMessages(channel chan string, exit chan bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ready to accept messages!!")

	for {
		input, _ := reader.ReadString('\n')
		if input == ExitText {
			break
		} else {
			channel <- input
		}
	}
	exit <- true
}
func main() {
	channel := make(chan string)
	exit := make(chan bool, 1)

	go receiver(channel)

	go ReadMessages(channel, exit)

	fmt.Println("Reading....")
	// Waits for the message before proceeding
	<-exit
	fmt.Print("Thank you!")
}

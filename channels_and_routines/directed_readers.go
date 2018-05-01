package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	ExitText   = "exit\n"
	BufferSize = 5
)

var wg sync.WaitGroup

func HandleMessage(message string) {
	fmt.Println("Message received: ", message)
}

func receiver(connection <-chan string) {
	fmt.Println("Setting up receiver...")
	for {
		HandleMessage(<-connection)
	}
	fmt.Println("Done in receiver")
}

func ReadMessages(channel chan<- string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your messages..")

	for {
		if len(channel) < BufferSize {
			input, _ := reader.ReadString('\n')
			if input == ExitText {
				wg.Done()
				fmt.Println("Thank you!! Have a great day!!")
				break
			} else {
				channel <- input
			}
			fmt.Println("Waiting for messages..")
		} else {
			fmt.Println("Receiver is busy processing.!!")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	channel := make(chan string, BufferSize)

	wg.Add(1)
	fmt.Println("Welcome!")

	go receiver(channel)
	go ReadMessages(channel)

	for len(channel) > 0 {
		time.Sleep(time.Second)
	}

	wg.Wait()
	fmt.Println("Done handling messages")
}

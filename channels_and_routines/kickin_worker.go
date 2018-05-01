package main

import (
	"bufio"
	"fmt"
	"os"
)

func HandleMessage(message string) {
	fmt.Println("Processing: ", message)
}

func worker(channel chan string) {
	fmt.Println("Starting worker...")
	for msg := range channel {
		HandleMessage(msg)
	}
	fmt.Println("Done in worker")
}

func ReadMessages(channel chan string) {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		input, _ := reader.ReadString('\n')
		if input == "exit\n" {
			break
		} else {
			channel <- input
		}
	}
	close(channel)
}

func main() {
	channel := make(chan string, 10)

	fmt.Println("Feed me words..")

	ReadMessages(channel)

	worker(channel)
	fmt.Println("Thank you!")
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	ExitText = "exit\n"
)

func StartChannel(user1 chan string, user2 chan string) {
	for {
		select {
		case message := <-user1:
			fmt.Println("poller-1 picked :: ", message)
		case message := <-user2:
			fmt.Println("poller-2 picked :: ", message)
		}
	}
}

func ReadMessages(user1 chan string, user2 chan string, exit chan bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ready to accept messages!!")

	for {
		input, _ := reader.ReadString('\n')
		if input == ExitText {
			break
			fmt.Print("Thank you...")
			exit <- true
		} else {
			if rand.Intn(20)%2 == 1 {
				user1 <- input
			} else {
				user2 <- input
			}
		}
	}
	exit <- true
}

func main() {
	user1 := make(chan string)
	user2 := make(chan string)
	exit := make(chan bool, 1)

	go ReadMessages(user1, user2, exit)
	go StartChannel(user1, user2)

	<-exit
}

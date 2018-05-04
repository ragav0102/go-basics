package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, queue chan float32) {
	for message := range queue {
		fmt.Println("worker ", id, " working on ", message)
		time.Sleep(2 * time.Second)
		fmt.Println("worker ", id, " done working on ", message, ".. Result:", message*message)
	}
	wg.Done()
}

func main() {
	queue := make(chan float32, 50)

	for i := 0; i < 20; i++ {
		queue <- rand.Float32()
	}

	close(queue)

	for i := 0; i < 5; i++ {
		go worker(i, queue)
		wg.Add(1)
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"time"
)

func BurstLimiter(BurstLimit int, rateLimit int, LimiterChannel chan int) {

	for i := 0; i < BurstLimit; i++ {
		LimiterChannel <- 1
	}
	// Rate limiting happens here. Limiterchannel holds processing for inter-request time
	for range time.Tick(time.Duration(1000/rateLimit) * time.Millisecond) {
		LimiterChannel <- 1
	}
}

func RateLimiter(rateLimit int, BurstLimit int, requests chan string) {

	LimiterChannel := make(chan int, BurstLimit)

	go BurstLimiter(BurstLimit, rateLimit, LimiterChannel)

	for range time.Tick(100 * time.Millisecond) {
		fmt.Println("waiting ", len(requests))
		for request := range requests {
			fmt.Println(time.Now(), " :: Processing request ::", request)
			<-LimiterChannel
		}
	}
}

func RequestGenerator(requests chan string) {
	for {
		if len(requests) < 30 {
			requests <- "requested at " + time.Now().String()
		} else {
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var rateLimit, BurstLimit int = 1, 1
	requests := make(chan string, 30)

	fmt.Println("Enter rate limit(per second).. ")
	fmt.Scanln(&rateLimit)
	fmt.Println("Enter burst limit.. ")
	fmt.Scanln(&BurstLimit)

	go RateLimiter(rateLimit, BurstLimit, requests)

	RequestGenerator(requests)
}

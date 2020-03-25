package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 1)
	stopSenderChan := make(chan bool)
	ticker := time.NewTicker(time.Second)

	go func() {
	Loop:
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			case <-stopSenderChan:
				break Loop
			}
		}
		fmt.Println("End. [sender]")
	}()
	var sum int
	for e := range intChan {
		fmt.Printf("Received: %v\n", e)
		sum += e
		if sum > 10 {
			fmt.Printf("Got: %v\n", sum)
			stopSenderChan <- true
			break
		}
	}
	fmt.Println("End. [receiver]")
}

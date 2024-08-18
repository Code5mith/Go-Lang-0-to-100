package main

// lets talk about switch statment
// The select statement in Go is used to wait on multiple communication operations. It's similar
// to a switch statement, but instead of evaluating expressions, it waits for channels to be ready
// for communication.

import (
	"fmt"
	"time"
)

func main() {
	// Create two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutines to send data to channels
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "Message from c1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "Message from c2"
	}()

	// Select statement to receive data from channels
	select {
	case msg1 := <-c1:
		fmt.Println("Received from c1:", msg1)
	case msg2 := <-c2:
		fmt.Println("Received from c2:", msg2)
	default:
		fmt.Println("No message received within the timeout")
	}
}

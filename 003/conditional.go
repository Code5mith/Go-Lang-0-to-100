package main

import "fmt"

func main() {

	message_length := 10
	max_message_len := 20

	if message_length > max_message_len {
		fmt.Println("Message is longer than maximum length allowed!")
	} else {
		fmt.Println("Message sent!")
	}
}

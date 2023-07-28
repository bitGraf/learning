package main

import "fmt"

func main() {
	// create a channel by defning its type
	messages := make(chan string)

	// chan <- 'pushes' data onto the channel
	go func() { messages <- "ping 1" }()
	// func() { messages <- "ping" }() ERROR: this does NOT work

	go func() { messages <- "ping 2" }()

	// <- chan 'pulles' data off the channel
	// channels are be default blocking until both the sender and receiver are ready
	// this syncs the program by default
	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}

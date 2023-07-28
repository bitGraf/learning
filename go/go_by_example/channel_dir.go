package main

import "fmt"

// this channel is send-only
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings channel is receive-only, pongs is send-only
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

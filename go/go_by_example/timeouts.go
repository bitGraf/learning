package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1) // non0blocking channel
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second): // this IS a blocking channel, so if <-c1 doesn't return within 1 second, this gets chosen
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1) // non0blocking channel
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second): // longer timeout, so the goroutine can finish in time
		fmt.Println("timeout 2")
	}
}

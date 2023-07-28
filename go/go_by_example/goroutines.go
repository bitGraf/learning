package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine") // runs asynchronously

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second) // manually wait for threads to finish
	fmt.Println("done")
}

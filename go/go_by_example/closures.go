package main

import "fmt"

func intSeq() func() int {
	// returns a function of type: func() int
	// the value of i is 'closed over' in the returning function
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq() // this function RETURNS a function

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println("7! =", fact(7))

	// recursive closure
	/*
		fib := func(n int) int {
			if n < 2 {
				return n
			}

			return fib(n-1) + fib(n-2)
		}
		This does NOT work!: fib is not defined beforehand, so it is not 'enclosed',
		so inside the function scope it cannot be used.
	*/
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}
	//This DOES work!: fib is explicitly defined beforehand, so it IS 'enclosed',
	//Inside the function scope it CAN be used.

	fmt.Println("fib(7):", fib(7))
}

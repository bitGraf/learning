package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// you can declare variables in the scope of the if-else, and use them in the conditional too :O
	if num := 11; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// not valid: clauses MUST be surrounded by braces
	//if true
	//	fmt.Println("hmm")

	// ternary operator does NOT exist ;~;
}

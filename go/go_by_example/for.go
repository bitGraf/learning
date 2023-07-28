package main

import "fmt"

func main() {
	i := 1
	for i <= 3 { // for loop with 1 condition -> is just a while loop?
		fmt.Println(i)
		i = i + 1 // does NOT automatically increment!
	}
	//i = 1
	//while i <= 3 { // 'while' does not exist :/
	//	fmt.Println(i)
	//	i = i + 1
	//}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for j := 0; j < 4; j++ { // j can be reused: j is 'scoped' to the for loop
		fmt.Println(j)
	}

	// continue and break work as expected
	for { // infinite loop
		fmt.Println("infinite loop step")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

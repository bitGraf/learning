package main

import "fmt"

// same pointer syntax as c
// *int is a pointer type
// &var is a reference to the variable
// *ptr dereferences a pointer

func zeroval(ival int) {
	ival = 0
}

func zeroptr(ival *int) {
	*ival = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

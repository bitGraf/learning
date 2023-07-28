package main

import "fmt"

// functions syntax:
// func function_name(arg1 arg1_type, arg2 arg2_type, ...) return_type {}

func plus(a int, b int) int {
	return a + b
}

// if args have the same type, they can be groups like this
func plusplus(a, b, c int) int {
	return a + b + c
}

func test(a, b, c int, d, e, f float32) map[string]map[string]int {
	m_inner := make(map[string]int)
	// m_inner["inner"] = (d * e * f) + a + b + c // ERROR: cannot add floats and ints -> no implicit cast
	m_inner["inner"] = int(d*e*f) + a + b + c

	m_outer := make(map[string]map[string]int)
	m_outer["outer"] = m_inner

	return m_outer
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusplus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	// wacky test
	m := test(1, 2, 3, 4.5, 5.6, 0.1)
	fmt.Println("m:", m)
}

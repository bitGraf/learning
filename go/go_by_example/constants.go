package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	k_not_const := n + 1 // := syntax does not inherit const-ness
	fmt.Println(k_not_const)
	k_not_const++
	fmt.Println(k_not_const)

	const k_const = n + 1
	fmt.Println(k_const)
	//k_const++ // can't do this: k_const is const
	fmt.Println(k_const)
}

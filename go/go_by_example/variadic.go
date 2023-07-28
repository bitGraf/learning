package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// arrays CANNOT be used to pass in like variadic args
	// you NEED to slice it first? [:] takes a slice of the whole array...
	num_arr := [...]int{1, 2, 3, 4}
	sum(num_arr[:]...)

	// slices can also be used to pass in like variadic args
	num_slice := []int{1, 2, 3, 4}
	sum(num_slice...)
}

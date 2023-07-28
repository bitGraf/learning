package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // range can be used to iterate over arrays and slices
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	sum = 0
	for i := range nums {
		sum += nums[i]
	}
	fmt.Println("sum:", sum)

	// general syntax:
	// for idx, value := range values { }
	// idx:   the index of the array 'values'
	// value: values[idx]
	// either idx or value can be ignored (or BOTH)
	count := 0
	for range nums {
		count++
	}
	fmt.Println("count:", count)

	// range on map
	// for key,value := range map {}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// can also do:
	// for key := range kvs {}
	// for _,value := range kvs {}

	// range can be used to iterate over strings
	str := "go lang"
	fmt.Println("str[0]:", str[0])     // string can be accessed like arrays
	fmt.Println("str[0:2]:", str[0:2]) // string can be accessed like arrays
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["l1"] // if the key does NOT exist => zero-value is returned for the type
	fmt.Println("map[l1]:", v1)

	v2 := m["k1"] // if the key does NOT exist => zero-value is returned for the type2
	fmt.Println("map[k1]:", v2)

	fmt.Println("len(m):", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// 1. : functions can return multiple values
	// 2. : you can ignore returns with _ (like ~ in MATLAB)
	// 3. : gettting a value from a map returns 2 things:
	//		value,present = map[key]
	_, present := m["k2"]
	fmt.Println("is k2 in m:", present)

	n := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]string{"foo": "1", "bar": "2"}

	// both print the same, since string values are printed WITHOUT "s...
	fmt.Println("n:", n)
	fmt.Println("n2:", n2)
}

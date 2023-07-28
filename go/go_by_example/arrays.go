package main

import "fmt"

func main() {
	var a [5]int // fixed-size array. by default, this is zero-valued
	fmt.Println("empty:", a)

	a[4] = 100
	fmt.Println("a =", a)
	fmt.Println("a[4] =", a[4])
	//fmt.Println("a[5] =", a[5]) // this error is caught at compile-time (fixed size is known)

	fmt.Println("len(a) =", len(a))

	// implicit syntax
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b =", b)

	//c := int{1, 2, 3, 4, 5} // does not work!
	c := []int{1, 2, 3, 4, 5} // this works! need to indicate its an array, but the size can be inferred from the {...}
	fmt.Println("c =", c)
	c = []int{5, 4, 3, 2, 1}
	fmt.Println("c =", c)
	fmt.Println(len(c))
	c = []int{5, 4, 3, 2}
	fmt.Println("c =", c)
	fmt.Println(len(c))
	c = []int{5, 4, 3, 2, 1, 0} // arrays are NOT fixed size?
	//c = []float32{5.0, 4.0, 3.0, 2.0, 1.0} // arrays ARE fixed type.
	fmt.Println("c =", c)
	fmt.Println(len(c))

	// nevermind... []int{...} is NOT a fixed sized array, it can be resized
	d := [5]int{1, 2, 3, 4, 5} // this IS fixed size
	fmt.Println("d =", d)
	d = [5]int{5, 4, 3, 2} // specifying less than max, the rest are zeroed
	fmt.Println("d =", d)
	//d = [5]int{5, 4, 3, 2, 1, 0} // Error: specifying more than max, causes error
	//fmt.Println("d =", d)

	//d = [5]float32{1, 2, 3, 4, 5} // Error: still fixed-type

	// 2-D arrays (arrays of arrays...)
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D: ", twoD)
}

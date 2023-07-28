package main

import "fmt"

func main() {
	var s []string
	fmt.Println("unint:", s, s == nil, len(s) == 0)

	// an array with no size... is a slice :/

	s = make([]string, 3)
	fmt.Println("empty:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	//s[3] = "oob" // causes a segfault
	s = append(s, "oob")
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, "1", "abcd")
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	fmt.Println("s[2:5]:", s[2:5]) // s[2:5] does NOT include 5!
	// [s[2], s[3], s[4]]

	fmt.Println("s[:5]:", s[:5]) // s[:5] includes everything UP TO (but not including) 5!
	// [s[0], s[1], s[2], s[3], s[4]]
	fmt.Println("s[2:]:", s[2:]) // s[2:] includes everything AFTER (and including) 2!
	// [s[2], s[3], s[4], ...]

	// copy slice
	var c = make([]string, len(s)-2)
	fmt.Println("Pre copy: ", c)
	copy(c, s) // if len(c) < len(s), only copy the size of c from s
	fmt.Println("Post copy:", c)
}

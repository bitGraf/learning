package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {

	// NOTE: Println appends spaces between comma separated args
	//       Print   does NOT
	/*
		fmt.Println("Write", i, "as ")
		fmt.Print("Write ", i, " as ")

		produce the same result... weird
	*/

	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("time.Now() =", time.Now())
	fmt.Println("time.Now().Weekday =", time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: // catching multiple values at once
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//var t time.Time = time.Now() // explicit syntax
	t := time.Now()
	switch { // switch without an expression is just an if/else statement
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's after noon")
	default:
		fmt.Println("It's noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type: %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	whatAmI(int(1))
	whatAmI(int64(1)) // int != int64
	whatAmI(int32(1)) // int != int32, int and int32 are distinct types

	var int_type int = 1
	var int32_type int32 = 1
	var int64_type int64 = 1
	fmt.Printf("sizeof(int) == %d\n", unsafe.Sizeof(int_type))
	fmt.Printf("sizeof(int32) == %d\n", unsafe.Sizeof(int32_type))
	fmt.Printf("sizeof(int64) == %d\n", unsafe.Sizeof(int64_type))

	// when I run this, in is a 64-bit integer. Even though, int != int64
}

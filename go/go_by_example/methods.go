package main

import "fmt"

type rect struct {
	width, height int
}

// struct methods are defined similar to how they are in C++ when defined outside the class definition.
// i.e:  void struct::foo() {this->member = 0;}
// is similar to
//       func (this *struct) foo() {this.member = 0;}

//    v this part is called the 'receiver' type
func (r *rect) area() int {
	return r.width * r.height
}

// receiver type can be passed by value as well
func (r rect) perim() int {
	r.width++ // this ONLY affects the perim calc, since r is passed by-value
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	r.width = 12
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
	fmt.Println("area:", rp.area()) // same as previous area(), since perim() has value receiver
}

package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p // huh???
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Jon"))

	s := person{age: 50, name: "Sean"}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age) // . is used for pointers too (no ->)
	sp.age = 51
	fmt.Println(sp.age)

	// anonymous structs
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}

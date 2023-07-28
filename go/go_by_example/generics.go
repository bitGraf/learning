package main

import "fmt"

// Generics: (like templates)
// func func_name[type parameters](args) return_type {}

// this function takes a generic map of type map[K]V
// K and V are type parameters, and are given constraints
// K must be comparable: vars of type K can be compared with == and !=
//					     this is required for map keys.
// V can be anything. (any is an alias for interface{})
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// define a generic type: linked list

type List[Type any] struct {
	head, tail *element[Type]
}
type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T2]) Push(v T2) {
	if lst.tail == nil {
		lst.head = &element[T2]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T2]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m)) // infers the generic types of MapKeys
	// can be specified however:
	MapKeys[int, string](m)

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())
}

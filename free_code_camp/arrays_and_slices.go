package main

import (
	"fmt"
)

// Agenda
// -> Arrays
// 	-> Creation
// 	-> Built-in functions
// 	-> Working with arrays
// -> Slices
// 	-> Creation
// 	-> Built-in functions
// 	-> Working with slices

var tCreation bool = false
var tArrayofArray bool = false
var tArrayCopy bool = false
var tSlice bool = false
var tMake bool = false
var tResizableSlice bool = false
var tSliceAsStack bool = true

func main() {
	if tCreation {
		grades := [...]int{100, 85, 93}
		fmt.Printf("Grades: %v", grades)

		var students []string
		fmt.Printf("Students: %v\n", students)
		students[0] = "Dadool"
		fmt.Printf("Stundents: %v", students)
	}

	if tArrayofArray {
		var identityMatrix [3][3]int = [3][3]int{
			{1, 0, 0},
			{0, 1, 0},
			{0, 0, 1}}
		fmt.Println(identityMatrix)
	}

	if tArrayCopy {
		a := [...]int{1, 2, 3}
		// b := &a -> shallow copy
		b := a
		b[1] = 1
		fmt.Println(a)
		fmt.Println(b)
	}

	if tSlice {
		a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		b := a[:]
		c := a[3:]
		d := a[3:6]
		e := a[:6]

		fmt.Printf("%v %v %v\n", a, len(a), cap(a))
		fmt.Printf("%v %v %v\n", b, len(b), cap(b))
		fmt.Printf("%v %v %v\n", c, len(c), cap(c))
		fmt.Printf("%v %v %v\n", d, len(d), cap(d))
		fmt.Printf("%v %v %v\n", e, len(e), cap(e))
	}

	if tMake {
		a := make([]int, 3, 100)
		fmt.Println(a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n", cap(a))
	}

	if tResizableSlice {
		a := []int{}
		fmt.Printf("Value: %v\n", a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n", cap(a))
		a = append(a, 1)
		fmt.Printf("Value: %v\n", a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n", cap(a))
		// a = append(a, 2, 3, 4, 5)
		a = append(a, []int{2, 3, 4, 5}...) // spread
		fmt.Printf("Value: %v\n", a)
		fmt.Printf("Length: %v\n", len(a))
		fmt.Printf("Capacity: %v\n", cap(a))
	}

	if tSliceAsStack {
		a := []int{1, 2, 3, 4, 5}
		b := append(a[:2], a[3:]...)
		fmt.Println(b)
	}
}

/*
Summary
-> Arrays
	-> Collection of items with the same type
	-> Fixed size
	-> Declaration styles
		-> a := [3]int{1,2,3}
		-> a := [...]int{1,2,3}
		-> var a [3]int
		-> 'len' function returns the size of the array
		-> Copies refer to different underlying data
-> Slices
	-> Backed by array
	-> Creation styles
		-> Slice existing array or slice
		-> Literal style
		-> Via 'make' function
			-> a := make([]int, 10)
			-> a := make([]int, 10, 10)
		-> 'append' function to add elements to the slice
		-> Copies refer to the same underlying data
*/

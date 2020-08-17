package main

/*
Agenda
-> Creating pointers
-> Dereference pointers
-> The 'new' function
-> Working with 'nil'
-> Types with internal pointers
*/

import (
	"fmt"
)

var tSimple bool = false
var tArithmetic bool = false
var tStruct bool = false
var tStruct1 bool = false
var tArrayVsSlice bool = false
var tMap bool = true

func main() {
	if tSimple {
		var a int = 42
		var b *int = &a    // * means pointers to an int
		fmt.Println(a, *b) // * is dereference operator
		fmt.Println(a, *b)
		a = 27
		fmt.Println(a, *b)
		fmt.Println(&a, b)
		*b = 42
		fmt.Println(a, *b)
	}

	if tArithmetic {
		a := [3]int{1, 2, 3}
		b := &a[0]
		// c := &a[1]-4 	<- Invalid operations
		c := &a[1]
		fmt.Printf("%v %p %p", a, b, c)
	}

	if tStruct {
		var ms *myStruct
		ms = &myStruct{foo: 42}
		fmt.Println(ms)
	}

	if tStruct1 {
		var ms *myStruct
		fmt.Println(ms) // -> gives <nil>
		ms = new(myStruct)
		ms.foo = 13         // * has lower precedence than the . operator
		fmt.Println(ms.foo) // (*ms).foo or ms.foo - both work fine
	}

	if tArrayVsSlice {
		// arrays -> sharing arrays is sharing values
		// a := [3]int{1, 2, 3}
		// arrays -> sharing slices is sharing data
		a := []int{1, 2, 3}
		b := a
		fmt.Println(a, b)
		a[1] = 42
		fmt.Println(a, b)
	}

	if tMap {
		a := map[string]string{"foo": "bar", "baz": "buz"}
		b := a
		fmt.Println(a, b)
		a["foo"] = "qux"
		fmt.Println(a, b)
	}
}

type myStruct struct {
	foo int
}

/*
Summary
-> Create pointers to objects
	-> can use & operator if value type already exists
		-> ms := myStruct{foo:42}
		-> p := &ms
	-> use addressof operator before initializer
		-> &myStruct{foo:42}
	-> use the 'new' keyword
		-> can't initialize fields at the same time
-> Types with internal pointers
	-> All assignment operations in 'Go' are copy operations
*/

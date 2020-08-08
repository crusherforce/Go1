//	Agenda
//
//	-> Variable Declaration
//	-> Redeclaration and shadowing
//	-> Visibility
//	-> Naming Conventions
//	-> Type Conversions

package main

import (
	"fmt"
)

// invalid: non-declaration statement outside the function body
// k := 6
var k int = 6

// unused declation works outside function scope
var l int = 7

var (
	k1 float32 = 1
	k2 float64 = 1
)

func main() {
	//	var i int
	//	i = 6 // loop initialization
	var j float64 = 7
	//	fmt.Println(i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", (float64)k1/k2, k1/k2)
}

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
	"math/rand"
	"strconv"
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

// Naming Convention and Scope of variable
// -> package level (lowercase) - package (scope)
// -> package level (uppercase) - exported (scope)
// -> block level - block level (scope)

var tInitialization bool = false
var tShortVarDeclaration bool = false
var tPointers bool = false
var tShodowing bool = false
var tConversion bool = true

func main() {
	if tInitialization {
		//	var i int
		//	i = 6 // loop initialization
		var j float64 = 7
		//	fmt.Println(i)
		fmt.Printf("%v, %T\n", j, j)
		fmt.Printf("%v, %T\n", k, k)

		var b, f, s = true, 2.3, "four"
		fmt.Printf("%T\t%T\t%T\n", b, f, s)
	}

	if tShortVarDeclaration {
		// short variable declaration -> name := expression
		freq := rand.Float64() * 3.0
		fmt.Printf("Value=%v, Type=%T\n", freq, freq)
	}

	if tPointers {
		// pointers
		x1 := 1
		p := &x1
		fmt.Println(*p)
		*p = 2
		fmt.Println(x1)

		// pointers (part 2)
		var x, y int
		fmt.Println(&x == &x, &x == &y, &x == nil)
	}

	if tShodowing {
		// shadowing, same var is declared outside the function's scope
		k1 := 25
		fmt.Println(k1)
	}

	if tConversion {
		var a1 = 466
		fmt.Printf("%v %T\n", a1, a1)

		var b1 float32
		b1 = float32(a1)
		fmt.Printf("%v %T\n", b1, b1)

		// gotcha
		var c1 string = strconv.Itoa(a1)
		fmt.Printf("%v %T\n", c1, c1)
	}

	/*
		Summary
		-> Variable declation (3 types)
		-> Can't redeclare variables, but can shadow them
		-> All variables must be used
		-> Visibility
			-> lower case first letter for package scope
			-> upper case first letter to export
			-> no private scope
		-> Naming Conventions
			-> Pascal/camel case
			-> As short as reasonable
		-> Type Conversions
			-> destination_type(variable)
			-> use strconv package for strings
	*/
}

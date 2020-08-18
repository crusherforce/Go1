package main

import "fmt"

var tBoolean bool = false
var tNumericTypes bool = false
var tComplex bool = false
var tString bool = false
var tRune bool = true

func main() {
	if tBoolean {
		n := 1 == 1
		m := 1 == 2
		fmt.Printf("%v %T\n", n, n)
		fmt.Printf("%v %T\n", m, m)
	}

	if tNumericTypes {
		var n uint64 = 42
		fmt.Printf("%v %T\n", n, n)

		a := 10
		b := 3
		fmt.Printf("%v %T\n", a+b, a+b)
		fmt.Printf("%v %T\n", a-b, a-b)
		fmt.Printf("%v %T\n", a/b, a/b)
		fmt.Printf("%v %T\n", a*b, a*b)
		fmt.Printf("%v %T\n", a%b, a%b)
	}

	if tComplex {
		var a complex64 = 1 + 2i
		var b complex64 = 2 + 5i
		fmt.Printf("%v %T\n", a+b, a+b)
		fmt.Printf("%v %T\n", a-b, a-b)
		fmt.Printf("%v %T\n", a/b, a/b)
		fmt.Printf("%v %T\n", a*b, a*b)
	}

	if tString {
		s := "this is a string"
		b := []byte(s)
		fmt.Printf("%v %T\n", b, b)
	}

	if tRune {
		r := 'a'
		fmt.Printf("%v %T\n", r, r)
	}
}

/*
Summary
-> Boolean type
	-> Values are true or false
	-> its not a type alias
-> Numeric types
	-> Integers
		-> Signed integers
			-> 'int' type has varying size, but min 32 bits
			-> 8 bit (int8) through 64 bit (int64)
		-> Unsigned integers
			-> 8 bit ('byte' or 'uint8') through 64 (uint64)
		-> Arithmetic operations
		-> Bitwise operators
		-> Can't mix type of different types
	-> Floating Point Numbers
		-> iEEE 754 standard
		-> Literal styles
			-> Decimal (3.14)
			-> Exponential (13e18 or 2E10)
			-> Mixed (13.7e12)
	-> Complex Numbers
		-> Zero value is (0+0i)
		-> 64 and 128 versions
		-> functions available
			-> complex, real, imag
	-> Strings
		-> UTF-8
		-> Immutable
		-> Can be concatenated with (+) operator
		-> Can be converted to []byte
	-> Rune
		-> UTF-32
		-> Alias for int-32
		-> Special methods normally required to process

*/

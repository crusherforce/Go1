package main

import (
	"fmt"
)

//	--- Agenda ---
//	-> Naming Convention
//	-> Typed Constants
//	-> Untyped Constants
//	-> Enumerated Constants
//	-> Enumerated Expressions

var tConst bool = false
var tOpWithVar bool = false
var tEnumConst1 bool = false
var tEnumConst2 bool = true

const a int16 = 27
const (
	_ = iota + 5
	cat_specialist
	dog_specialist
	snake_specialist
)
const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
const (
	isAdmin = 1 << iota
	isHQ
	canSeeFinance

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNA
	canSeeSA
)

func main() {
	if tConst {
		const a uint64 = uint64(1)
		const b string = "baby teddy"
		const c float32 = 3.14
		const d bool = true

		fmt.Printf("%v %T\n", a, a)
		fmt.Printf("%v %T\n", b, b)
		fmt.Printf("%v %T\n", c, c)
		fmt.Printf("%v %T\n", d, d)
	}

	if tOpWithVar {
		// const a int = 42 (Invalid)
		const a = 42
		var b int16 = 27
		fmt.Printf("%v %T\n", a+b, a+b)
	}

	if tEnumConst1 {
		fmt.Printf("%v %T\n", cat_specialist, cat_specialist)
		fmt.Printf("%v %T\n", dog_specialist, dog_specialist)
		fmt.Printf("%v %T\n", snake_specialist, snake_specialist)

		fmt.Printf("%v %T\n", KB, KB)
		fmt.Printf("%v %T\n", MB, MB)
		fmt.Printf("%v %T\n", GB, GB)
		fmt.Printf("%v %T\n", TB, TB)
		fmt.Printf("%v %T\n", PB, PB)
		fmt.Printf("%v %T\n", EB, EB)
	}

	if tEnumConst2 {
		var roles byte = isAdmin | canSeeFinance | canSeeEurope
		fmt.Printf("%b\n", roles)
		fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
		fmt.Printf("Is HQ? %v\n", isHQ&roles == isHQ)
	}
}

/*
-- Summary
	-> Immutable but can be shadowed
	-> Replaced by compiler at compile time
	-> Named like variables
	-> Typed constants work like immutable variables
	-> Untyped constants work like literals
	-> Enumerated Constants
		-> 'iota' allows related constants to be created easily
		-> 'iota' starts at 0 in each const block and increments by one
		-> Watch out of constant values that match zero val of variable
	-> Enumerated Expressions
		-> Operations that can be determited at compile time are allowed
			-> Arithmetic
			-> Bitwise ops
			-> Bitshifting
*/

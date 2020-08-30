package main

import (
	"fmt"
)

var tBitwise = true

func main() {
	if tBitwise {
		var x uint8 = 1<<5 | 1<<1
		var y uint8 = 1<<2 | 1<<1
		fmt.Printf("%08b\n", x) // Prints 00100010
		fmt.Printf("%08b\n", y) // Prints 00000110

		// appying some bitwise opearators
		fmt.Printf("%08b\n", x&y)  // intersection
		fmt.Printf("%08b\n", x|y)  // union
		fmt.Printf("%08b\n", x^y)  // symmetric difference
		fmt.Printf("%08b\n", x&^y) // difference

		for i := uint(0); i < 8; i++ {
			if x&(1<<i) != 0 {
				fmt.Println(i)
			}
		}
	}
}

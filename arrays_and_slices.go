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

var tCreation bool = true

func main() {
	if tCreation {
		grades := [...]int{97, 85, 93}
		fmt.Printf("Grades: %v", grades)

		var students []string
		fmt.Printf("Students: %v\n", students)
		students[0] = "Dadool"
		fmt.Printf("Stundents: %v", students)
	}
}

package main

import "fmt"

const freezingF, boilingF = 32.0, 212.0

func main() {
	fmt.Printf("freezing point = %g°F or %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("boiling point = %g°F or %g°C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

// A function declaration has -
// -> a name
// -> a list of parameters
// -> an optional list of results
// -> the function body

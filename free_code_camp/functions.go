package main

/*
	 Agenda
	===========
	-> Basic Syntax
	-> Parameters
	-> Return values
	-> Anonymus functions
	-> Functions as types
	-> Methods
*/
import (
	"fmt"
)

var tSimple bool = false
var tVariatic bool = false
var tIdiomatic bool = false
var tAnon bool = false
var tFuncVar bool = false
var tMethod bool = true

func main() {
	if tSimple {
		ms1, ms2 := "hello", "world"
		sayMessage(ms1, ms2)
		fmt.Println(ms1, ms2)
		sayMessagePtr(&ms1, &ms2)
		fmt.Println(ms1, ms2)
	}

	if tVariatic {
		sum("The sum is", 1, 2, 3, 4, 5)
		s := mult(1, 2, 3, 4, 5)
		fmt.Println("The product is", s)
	}

	if tIdiomatic {
		d, err := divide(5.0, 3.0)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(d)
		}
	}

	if tAnon {
		f := func() {
			fmt.Println("Hello Go!")
		}
		f()
	}

	if tFuncVar {
		// Invalid
		// d, err := divide(5.0, 3.0)

		var divide func(float64, float64) (float64, error)
		divide = func(a, b float64) (float64, error) {
			if b == 0.0 {
				return 0.0, fmt.Errorf("Can't divide by zero")
			} else {
				return a / b, nil
			}
		}

		d, err := divide(5.0, 3.0)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(d)
		}
	}

	if tMethod {
		g := greeter{
			greeting: "Hello",
			name:     "Go",
		}
		g.greet()
	}
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func (g *greeter) greetPtr() {
	fmt.Println(g.greeting, g.name)
}

func divide(a float64, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func sum(msg string, values ...int) {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println(msg, result)
}

func mult(values ...int) (result int) {
	fmt.Println(values)
	result = 1
	for _, v := range values {
		result *= v
	}
	return
}

func sayMessage(ms1, ms2 string) { // cool stuff
	fmt.Println(ms1, ms2)
	ms2 = "earth" // has no effect, as the params are passed by value
}

func sayMessagePtr(ms1, ms2 *string) {
	fmt.Println(*ms1, *ms2)
	*ms2 = "earth" // has effect
}

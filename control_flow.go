package main

import "fmt"

var tIf bool = false
var tSwitch bool = false
var tTypeSwitch bool = true

func main() {
	if tIf {
		statePopulations := map[string]int{"West Bengal": 100000000, "Uttar Pradesh": 200000000}
		if pop, ok := statePopulations["Karnataka"]; ok {
			fmt.Println(pop)
		}
	}

	if tSwitch {
		var v int = 3
		switch {
		case v <= 10:
			fmt.Println("less than or equal to 10")
			fallthrough
		case v <= 20:
			fmt.Println("less than or equal to 20")
			fallthrough
		default:
			fmt.Println("all else")
		}
	}

	if tTypeSwitch {
		type Doctor struct {
			number     int
			actorName  string
			companions []string
		}
		var i interface{} = new(Doctor)
		switch i.(type) {
		case int16:
			fmt.Println("it's int16")
		case int:
			fmt.Println("it's int")
		case float32:
			fmt.Println("it's float32")
		case float64:
			fmt.Println("it's float64")
		case []int:
			fmt.Println("it's []int")
		case Doctor:
			fmt.Println("it's Doctor")
		default:
			fmt.Println("it's other type")
		}
		fmt.Printf("%v %T\n", i, i)
	}
}

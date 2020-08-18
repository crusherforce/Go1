package main

import "fmt"

func main() {
	tLoop1 := false
	tLoop2 := false
	tLoop3 := false
	tLoopRange := true

	if tLoop1 {
		for i, j := 0, 0; i*j <= 25; i, j = i+1, j+1 {
			fmt.Println(i, j)
		}
	}

	if tLoop2 {
		i := 0
		for i < 5 {
			fmt.Println(i)
			i++
		}
	}

	if tLoop3 {
		i := 0
		for {
			fmt.Println(i)
			if i >= 5 {
				break
			}
			i++
		}
	}

	if tLoopRange {
		s := []int{1, 2, 3}
		for k, v := range s {
			fmt.Println(k, v)
		}

		statePopulations := map[string]int{"West Bengal": 100000000, "Uttar Pradesh": 200000000}
		for k, _ := range statePopulations {
			fmt.Println(k)
		}
	}
}

/*
Summary:
	-> 'for' statement
		-> simple loops
			-> for initializer; test; incrementer {}
			-> for test {}
			-> for {}
		-> exiting early
			-> break
			-> continue
			-> labels
		-> looping over collections
			-> arrays, slices, string, maps, channels
			-> for k,v := range collection {}
*/

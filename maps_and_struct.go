package main

import (
	"fmt"
	"reflect"
)

var tMaps bool = false
var tStruct bool = false
var tAnonStruct bool = false
var tEmbedding bool = false
var tTags bool = true

func main() {
	if tMaps {
		statePopulations := make(map[string]int)
		statePopulations = map[string]int{"West Bengal": 100000000, "Uttar Pradesh": 200000000}
		var ok bool
		_, ok = statePopulations["West Bengal"]
		fmt.Println(ok)
		_, ok = statePopulations["Random"]
		fmt.Println(ok)
		fmt.Println(len(statePopulations))
		sp := statePopulations
		delete(sp, "West Bengal")
		fmt.Println(sp)
		fmt.Println(statePopulations)
	}

	if tStruct {
		type Doctor struct {
			number     int
			actorName  string
			companions []string
		}

		aDoctor := Doctor{
			number:    3,
			actorName: "Jon Pertwee",
			companions: []string{
				"Liz Shaw",
				"Jo Grant"}}

		fmt.Println(aDoctor.companions[1])
	}

	if tAnonStruct {
		aDoctor := struct{ name string }{name: "Mr. Random"}
		anotherDoctor := aDoctor
		anotherDoctor.name = "Dr. Random"
		fmt.Println(aDoctor)
		fmt.Println(anotherDoctor)
	}

	if tEmbedding {
		type Animal struct {
			Name   string
			Origin string
		}

		type Bird struct {
			Animal
			SpeedKPH float32
			CanFly   bool
		}

		b := Bird{
			Animal: Animal{
				Name:   "Emu",
				Origin: "Australia"},
			SpeedKPH: 48,
			CanFly:   false,
		}

		fmt.Println(b)
	}

	if tTags {
		type Animal struct {
			Name   string `required max:"100"`
			Origin string
		}

		t := reflect.TypeOf(Animal{})
		field, _ := t.FieldByName("Name")
		fmt.Println(field)
	}
}

/*
Summary:
-> Maps
	-> Collections of value types that are accessed via keys
	-> Created via literals or via 'make' functions
	-> Members accessed via myMap["key"]
	-> Check for presence with "value,ok" form of result
-> Structs
	-> Collections of disparate data types
	-> Keyed by named fields
	-> Normally created as types, but anon structs are allowed
	-> Structs are value types
	-> No inheritance, but can use composition via embedding
	-> Tags can be added to struct fields to describe field
*/

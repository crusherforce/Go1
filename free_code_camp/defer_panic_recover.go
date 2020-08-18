package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var tSimple bool = false
var tHttp bool = false
var tUnique bool = false
var tPanic bool = false
var tPanicAgain bool = false
var tPanicDobara bool = false
var tOrderOfExecution bool = false
var tOrderOfExecution1 bool = false
var tPanicker bool = true

func main() {
	if tSimple {
		defer fmt.Println("start")
		defer fmt.Println("middle")
		defer fmt.Println("end")
	}

	if tHttp {
		res, err := http.Get("http://www.google.com/robots.txt")
		if err != nil {
			log.Fatal(err)
		}
		// most practical use of defer: closing resources immediately
		// after referencing it
		defer res.Body.Close()
		robots, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", robots)
	}

	if tUnique {
		a := "start"
		defer fmt.Println(a)
		a = "end"
		// spoiler: it prints "start"
	}

	if tPanic {
		a, b := 1, 0
		ans := a / b
		fmt.Println(ans)
	}

	if tPanicAgain {
		fmt.Println("start")
		panic("Something happened")
		fmt.Println("end")
	}

	if tPanicDobara {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Go!"))
		})
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err.Error())
		}
	}

	if tOrderOfExecution {
		// main() -> defer -> panic -> return
		fmt.Println("Start")
		defer fmt.Println("This was deferred")
		panic("Something bad happened")
		fmt.Println("End")
	}

	if tOrderOfExecution1 {
		fmt.Println("Start")
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error: ", err)
			}
		}()
		panic("Something bad happened")
		fmt.Println("End")
	}

	if tPanicker {
		fmt.Println("Start")
		panicker()
		fmt.Println("End")
	}
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("defer panicking")
}

/*
Summary:
->defer
	-> delay execution of statement until the end of the function
	-> useful to group 'open' and 'close' functions together
		-> beware of loops
	-> multiple defer statements : LIFO order
-> panic
	-> occurs when program can't continue at all
		-> don't use for file can't be opened, return error
		-> unrecoverable events
	-> functions will stop executing
		-> deferred function will still fire
	-> if nothing handles panic, program will exit
-> recover
	-> used to recover from panics
	-> only useful in deferred functions
	-> current function will not attempt to continue, but higher functions in the call stack will
*/

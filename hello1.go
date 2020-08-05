package main

import (
	"fmt"
	"time"
	"net/http"
	"hash/crc32"
	"encoding/hex"
	"os"
	"reflect"
)

func main() {
	fmt.Printf("Hello %s\n", new(World))

	// Printing Durations
	start := time.Now()

	// Make an HTTP request
	resp, _ := http.Get("http://www.google.com")
	resp.Body.Close()
	fmt.Println(time.Since(start))

	h := crc32.NewIEEE()
	fmt.Fprintf(h, "Hello World")
	fmt.Printf("hash=%x\n", h.Sum32())

	h1 := hex.Dumper(os.Stdout)
	// defer h1.Close()
	// fmt.Fprintf(h1, "Hello World\n")

	type Lang struct {
		Name string
		Year int
		URL string	
	}
	
	lang := Lang{"Go", 2009, "http://golang.org"}
	fmt.Printf("%v\n", lang) // or %+v
}

type World struct{}

func (w *World) String() string {
	return "Crusherforce"
}

func myPrint(args ...interface{}) {
	for _, arg := range args {
		switch v := reflect.ValueOf(arg); v.Kind() {
		case reflect.String:
			os.Stdout.WriteString(v.String())
		case reflect.Int:
			os.Stdout.WriteString(strconv.FormatInt(v.Int(), 10))
		}
	}
}

/*
Notes
------
Ref: A Tour of Go (https://research.swtch.com/gotour)

- A type implements an interface by defining the required methods
- Reflection : Type information, basic opearations available at run-time.
*/
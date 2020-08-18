package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"hash/crc32"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"
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

	// h1 := hex.Dumper(os.Stdout)
	// defer h1.Close()
	// fmt.Fprintf(h1, "Hello World\n")

	type Lang struct {
		Name string
		Year int
		URL  string
	}

	lang := Lang{"Go", 2009, "http://golang.org"}
	fmt.Printf("%v\n", lang)
	fmt.Printf("%+v\n", lang)
	fmt.Printf("%#v\n", lang)

	data, err := json.Marshal(lang)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)

	data1, err1 := xml.MarshalIndent(lang, "", " ")
	if err1 != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data1)
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

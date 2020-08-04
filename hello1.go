package main

import (
	"fmt"
	"time"
	"net/http"
	"hash/crc32"
	"encoding/hex"
	"os"
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
	defer h1.Close()
	fmt.Fprintf(h1, "Hello World\n")
}

type World struct{}

func (w *World) String() string {
	return "Crusherforce"
}

/*
Notes
------
Ref: A Tour of Go (https://research.swtch.com/gotour)

- A type implements an interface by defining the required methods
*/
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	do(func(lang Lang) {
		fmt.Printf("%v\n", lang)
	})

	do(func(lang Lang) {
		data, err := xml.MarshalIndent(lang, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", data)
	})

	start := time.Now()
	c := make(chan string)
	n := 0
	do(func(lang Lang) {
		n++
		go count(lang.Name, lang.URL, c)
	})

	timeout := time.After(10 * time.Second)
	for i := 0; i < n; i++ {
		select {
		case result := <-c:
			fmt.Print(result)
		case <-timeout:
			fmt.Print("Timed out\n")
			return
		}
	}

	fmt.Printf("%.2fs total\n", time.Since(start).Seconds())
}

func count(name, url string, c chan<- string) {
	start := time.Now()
	r, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("%s: %s\n", name, err)
		return
	}
	n, _ := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	c <- fmt.Sprintf("%s %d [%.2fs]\n", name, n, time.Since(start).Seconds())
}

func do(f func(Lang)) {
	input, err := os.Open("lang.json")
	if err != nil {
		log.Fatal(err)
	}
	// io.Copy(os.Stdout, input)
	dec := json.NewDecoder(input)
	for {
		var lang Lang
		err := dec.Decode(&lang)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		f(lang)
	}
}

type Lang struct {
	Name string
	Year int
	URL  string
}

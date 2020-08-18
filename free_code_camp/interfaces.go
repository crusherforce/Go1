package main

/*
Agenda(Interfaces)
==================
-> Basics
-> Composing Interfaces
-> Type Conversion
	-> The empty interface
	-> Type switches
-> Implementing with values vs pointers
-> Best practices
*/

import (
	"bytes"
	"fmt"
)

var tWriter bool = false
var tIncrementer bool = false
var tComposeInterface bool = false
var tEmptyInterface bool = true

func main() {
	if tWriter {
		var w Writer = ConsoleWriter{}
		w.Write([]byte("Hello Go!"))
	}

	if tIncrementer {
		myInt := IntCounter(0)
		var inc Incrementer = &myInt
		for i := 0; i < 10; i++ {
			fmt.Println(inc.Increment())
		}
	}

	if tComposeInterface {
		var wc WriterCloser = CreateBufferedWriterCloser()
		wc.Write([]byte("Hello there, my friend!"))
		wc.Close()

		bwc, ok := wc.(*BufferedWriterCloser)
		if ok {
			fmt.Println(bwc)
		} else {
			fmt.Println("Type Conversion Failed.")
		}
	}

	if tEmptyInterface {
		var myObj interface{} = CreateBufferedWriterCloser()
		// Casting empty interface to a WriterCloser
		if wc, ok := myObj.(WriterCloser); ok {
			wc.Write([]byte("Hello world once again."))
			wc.Close()
		}
	}
}

func CreateBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}
	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

/*
	Best Practices
	==============
	-> Use many, small interfaces
		-> Single method interfaces are most powerful and flxeible
			-> e.g. - io.Writer, io.Reader, interface{}
	-> Don't export interfaces for types that will be consumned
	-> Do export interfaces for types that will be used by the package
	-> Design functions and methods to receive interfaces
*/

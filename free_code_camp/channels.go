package main

/*
	Agenda
	======
	-> Channel Basics
	-> Restricting Data Flow
	-> Buffered Channels
	-> Closing Channels
	-> For ... range loops with channels
	-> nested statements
*/

import (
	"fmt"
	"sync"
)

var tSimple = false
var tSendReceive = false
var tLoopRange = false
var tAnotherGoPresentation = true

var wg = sync.WaitGroup{}

func main() {
	if tSimple {
		ch := make(chan int) // channel is strongly typed
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)
			ch <- 27
			wg.Done()
		}()
		go func() {
			ch <- 42
			fmt.Println(<-ch)
			wg.Done()
		}()
		wg.Wait()
	}

	if tSendReceive {
		ch := make(chan int, 50) // buffered-channel
		wg.Add(2)
		go func(ch <-chan int) { // receive-only channel
			i := <-ch
			fmt.Println(i)
			wg.Done()
		}(ch)
		go func(ch chan<- int) { // send-only channel
			ch <- 42
			ch <- 27
			wg.Done()
		}(ch)
		wg.Wait()
	}

	if tLoopRange {
		ch := make(chan int, 50)
		wg.Add(2)
		go func(ch <-chan int) {
			// for i := range ch {
			// 	fmt.Println(i)
			// }
			for {
				if i, ok := <-ch; ok {
					fmt.Println(i)
				} else {
					break
				}
			}
			wg.Done()
		}(ch)
		go func(ch chan<- int) {
			ch <- 42
			ch <- 27
			close(ch)
			wg.Done()
		}(ch)
		wg.Wait()
	}

	if tAnotherGoPresentation {
		ch := computeAndSend(x,y,z)
		v2 := anotherExpensiveComputation(a,b,c)
		v1 := <-ch
		fmt.Println(v1,v2)
	}
}

func computeAndSend(x,y,z int) chan int {
	ch := make(chan, int)
	go func() {
		ch <- expensiveComputation(x,y,z)
	}()
	return ch
}

func expensiveComputation(x,y,z int) {
}

func anotherExpensiveComputation(x,y,z int) {	
}
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 	Agenda
// 	=======
//	-> Creating goroutines
//	-> Synchronization
//		-> Waitgroups
//		-> Mutexes
//	-> Parallelism
//	-> Best Practices

var tBasic = false
var tWaitgroup = false
var tThreadsAvail = false
var tSynchronization = false
var tCompileTimeRaceCheck = true

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	if tBasic {
		// Observed 3 different execution sequences
		// -> 1
		// Inside anon2(), Hellllo, Inside sayHello(), Hello, Inside anon1(), Hello
		// -> 2
		// Inside anon2(), Hellllo, Inside anon1(), Hello, Inside sayHello(), Hello
		// -> 3
		// Inside sayHello(), Hello, Inside anon1(), Inside anon2(), Hellllo, Hello

		msg := "Hello"
		// prints "Hello"
		go sayHello()
		// prints "Hello"
		go func(msg string) {
			fmt.Println("Inside anon1()")
			fmt.Println(msg)
		}(msg)
		// prints "Hellllo"
		go func() {
			fmt.Println("Inside anon2()")
			fmt.Println(msg)
		}()
		msg = "Hellllo"
		time.Sleep(5 * time.Millisecond)
	}

	if tWaitgroup {
		var msg = "Hello"
		wg.Add(1)
		go func(msg string) {
			fmt.Println(msg)
			wg.Done()
		}(msg)
		msg = "Goodbye"
		wg.Wait()
	}

	if tThreadsAvail {
		fmt.Printf("OS threads : %v\n", runtime.GOMAXPROCS(-1))
		runtime.GOMAXPROCS(2)
		fmt.Printf("(Modified) OS threads : %v\n", runtime.GOMAXPROCS(-1))
	}

	if tSynchronization {
		runtime.GOMAXPROCS(100)
		for i := 0; i < 10; i++ {
			wg.Add(2)
			m.RLock()
			go sayHello()
			m.Lock()
			go increment()
		}
		wg.Wait()
	}

	if tCompileTimeRaceCheck {
		// cmd:	go run -race Go1/goroutines.go
		// ==================
		// WARNING: DATA RACE
		// Read at 0x00c0000941e0 by goroutine 7:
		//   main.main.func4()
		//       C:/Users/meroycha/Documents/Learning/Go/Go1/goroutines.go:107 +0x43

		// Previous write at 0x00c0000941e0 by main goroutine:
		//   main.main()
		//       C:/Users/meroycha/Documents/Learning/Go/Go1/goroutines.go:109 +0x1f3

		// Goroutine 7 (running) created at:
		//   main.main()
		//       C:/Users/meroycha/Documents/Learning/Go/Go1/goroutines.go:106 +0x1e5
		// ==================
		// Value2
		// Found 1 data race(s)
		// exit status 66

		var msg = "Value1"
		go func() {
			fmt.Println(msg)
		}()
		msg = "Value2"
		time.Sleep(5 * time.Millisecond)
	}
}

func sayHello() {
	fmt.Printf("Hello %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

// 	Best Practices
// 	==============
// 	-> Don't create goroutines in libraries
// 		-> Let consumer control concurrency
// 	-> When creating a goroutine, know how it will end
//		-> Avoids subtle memory leaks
// 	-> Check for race conditions at compile time
//

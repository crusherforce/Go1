package main

import (
	"fmt"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{})

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is Starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(1000 * time.Millisecond)
}

func logger() {
	for {
		select {
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2020-08-19"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}

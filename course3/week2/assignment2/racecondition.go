// Following code shows an example of a race condition in Go.
//
// We have two go routines that access `message` variable:
// - Goroutine 1 - sets the `message` to "Hello"
// - Goroutine 2 - prints the `message`
//
// Race condition occurs because the order of execution is non-deterministic.
// There is no synchronisation mechanism, so we might print `message` before it's been set.
//
// Please run following program a couple of times, and check the output.
// Program should sometimes print "Message:", where the other time it is "Message: Hello".

package main

import (
	"fmt"
	"time"
)

func main() {
	var message string

	// Goroutine 1 sets the message to "Hello"
	go func() {
		message = "Hello"
	}()

	// Goroutine 2 prints the message
	go func() {
		fmt.Println("Message:", message)
	}()

	time.Sleep(time.Second) // Wait for goroutines to finish
}

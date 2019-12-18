package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = "Hello"
	go func() {
		fmt.Println(msg) // this line makes this goroutine dependent on the main thread
		// this will print "Goodbye"
	}()

	// generally, the go scheduler will not interrupt the main function unless it hit the time.Sleep() function

	go func(msg string) {
		fmt.Println(msg)
	}(msg) // by using a parameter, we decoupled the msg variable and the main function

	msg = "Goodbye"
	time.Sleep(1000 * time.Millisecond)
}

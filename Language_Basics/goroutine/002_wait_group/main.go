package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // waitGroup is designed to synchronize multiple goroutines together
// we want to synchronize our goroutine to the main function's goroutine

func main() {
	var msg = "hello"

	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	msg = "Goodbye"

	wg.Wait() // this line makes sure that the main function will not exit out too early
}

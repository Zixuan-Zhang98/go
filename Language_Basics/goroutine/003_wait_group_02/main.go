package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2) // here, our goroutines are actually raising against each other
		// there's no synchronization between these goroutines and they just run as fast as they can
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println(counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}

package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{} // this protects data that only one thing can access it at a time

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(2) // here, our goroutines are actually raising against each other
		// there's no synchronization between these goroutines and they just run as fast as they can
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println(counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

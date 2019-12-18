package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(1000 * time.Millisecond)
}

func sayHello() {
	fmt.Println("hello world")
}

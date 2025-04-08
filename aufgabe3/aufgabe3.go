package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hallo von der Goroutine!")
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)
	fmt.Println("Hallo von der Main!")
}

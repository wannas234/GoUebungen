package main

import (
	"fmt"
	"time"
)

// Aufgabe1.3: Erstelle eine Go Routine
func sayHello() {
	fmt.Println("Hallo von der Goroutine!")
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)
	fmt.Println("Hallo von der Main!")
}

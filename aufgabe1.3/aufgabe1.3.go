package main

import (
	"fmt"
	"time"
)

// Aufgabe1.3: Erstelle eine Go Routine
func GoRoutine() {
	fmt.Println("Hallo von der Goroutine!")
}

func main() {
	go GoRoutine()
	time.Sleep(1 * time.Second)
	fmt.Println("Hallo von der Main!")
}

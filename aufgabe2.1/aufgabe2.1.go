package main

import "fmt"

// Aufgabe2.1: Erstelle einen Channel
func main() {

	channel := make(chan string)

	go func() {
		channel <- "Hallo, Welt!"
	}()

	message := <-channel
	fmt.Println(message)
}

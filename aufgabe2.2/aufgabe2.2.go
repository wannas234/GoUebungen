package main

import (
	"fmt"
	"time"
)

// Aufgabe2.2: Erstelle mithilfe von Go Routinen und Channels einen Dialog zwischen Lisa und Manfred
func main() {

	messageChannel := make(chan string)

	go func() {
		messageChannel <- "Lisa: Guten Tag"
		time.Sleep(1 * time.Second)
		messageChannel <- "Lisa: Lang nichts mehr gehört!"
		time.Sleep(1 * time.Second)
		messageChannel <- "Lisa: Was machen Sachen?"
		time.Sleep(1 * time.Second)
		messageChannel <- "Manfred: Heyyyyy"
		time.Sleep(1 * time.Second)
		messageChannel <- "Manfred: Bei mir geht nichts, Was Geht bei dir??"
		close(messageChannel)
	}()

	for msg := range messageChannel {
		fmt.Println(msg)
	}
}

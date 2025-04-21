package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GO Channels Aufgabe für Mega Ultra Profis :Sonnenbrillenemoji:

// Aufgabenstellung: Erstelle ein ausführbares GO-Projekt mit einer Funktion namens
// "boring", welche eine Nachricht als Funktionsargument annimmt und diese permanent
// mit einem Delay zwischen 1 - 1000ms an einen Kanal sendet. Die entsandten
// Nachrichten sollen in der Main-Funktion abgefangen und in der
// Konsole ausgegeben werden.
// Nach der fünften Nachricht wird in der Konsole "Mir reichts. Du bist langweilig."
// ausgegeben und das Programm beendet.

// Tipp: Benutzt die Pakete "fmt", "time" und "math/rand"

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d\n", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("Langweilig!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("Du sagst %s", <-c)
	}
	fmt.Println("Mir reichts. Du bist langweilig.")
}

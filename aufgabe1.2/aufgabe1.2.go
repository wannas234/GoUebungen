package main

import (
	"fmt"
	"time"
)

// Aufgabe1.2: Kombiniere time.Sleep mit Println um einen Delay zu simulieren.
func main() {

	fmt.Println("Warte 3 Sekunden...")

	time.Sleep(3 * time.Second)

	fmt.Println("Hallo, Welt!")
}

package main

import (
	"fmt"
	"time"
)

// Aufgabe1.2: Kombiniere time.Sleep mit Println um nach einer Konsolen Ausgabe, die nächste Ausgabe erst nach 3 Sekunden starten zu lassen.
func main() {

	fmt.Println("Warte 3 Sekunden...")

	time.Sleep(3 * time.Second)

	fmt.Println("Hallo, Welt!")
}

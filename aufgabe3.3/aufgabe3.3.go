package main

import "fmt"

// Aufgabe3.3: Erstelle ein Programm, welches die ersten 6 Fibonacci-Zahlen ausgibt.

// Die Funktion `fib` gibt eine Funktion zurück, die die nächste Fibonacci-Zahl berechnet.
func fib() func() int {
	a, b := 0, 1 // Initialisiere die ersten beiden Fibonacci-Zahlen: a = 0, b = 1
	return func() int {
		a, b = b, a+b // Berechne die nächste Fibonacci-Zahl: a wird b, b wird a+b
		return a      // Gib die aktuelle Fibonacci-Zahl (a) zurück
	}
}

func main() {
	f := fib() // Erstelle eine Instanz der Fibonacci-Funktion
	// Rufe die Funktion `f` sechsmal auf, um die ersten 6 Fibonacci-Zahlen zu berechnen und auszugeben
	fmt.Println(f(), f(), f(), f(), f(), f())
}

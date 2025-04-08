package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Warte 3 Sekunden...")

	time.Sleep(3 * time.Second)

	fmt.Println("Hallo, Welt!")
}

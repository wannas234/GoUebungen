package main

import "fmt"

func fibonacci(n int) {

	a, b := 0, 1

	for a <= n {
		fmt.Print(a, " ")

		a, b = b, a+b
	}
}

func main() {
	fibonacci(35)
}

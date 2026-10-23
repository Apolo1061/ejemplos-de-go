package main

import (
	"fmt"
)

func printS(t func(int, int) int, a int, b int) {
	fmt.Printf("%d + %d = %d \n", a, b, t(a, b))
}

func printR(t func(int, int) int, a int, b int) {
	fmt.Printf("%d - %d = %d \n", a, b, t(a, b))
}

func main() {
	suma := func(a int, b int) int {
		return a + b
	}

	resta := func(a int, b int) int {
		return a - b
	}

	num1 := 2
	num2 := 2

	printS(suma, num1, num2)
	printR(resta, num1, num2)
}

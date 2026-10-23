package main

import (
	"fmt"
)

func sumador(a int, b int) int {
	return a + b
}

func main() {
	var num1 = 5
	var num2 = 3

	var res int = sumador(num1, num2)

	fmt.Printf("num1 + num2 = %d", res)

}

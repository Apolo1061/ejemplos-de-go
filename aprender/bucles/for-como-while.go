package main

import (
	"fmt"
)

func main() {
	var valor int = 0

	for valor < 5 {
		fmt.Printf("valor = %d \n", valor)
		valor++
	}
}

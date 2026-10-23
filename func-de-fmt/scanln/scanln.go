package main

import (
	"fmt"
)

func main() {
	var nombre string
	var edad int

	fmt.Printf("\nIngresa tu nombre: ")
	fmt.Scanln(&nombre)

	fmt.Printf("\nIngresa tu edad: ")
	fmt.Scanln(&edad)

	fmt.Println("te llamas: ", nombre, " y tu edad es: ", edad)
}

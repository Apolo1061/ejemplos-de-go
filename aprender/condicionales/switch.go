package main

import (
	"fmt"
)

func main() {
	var opt string
	fmt.Print("Ingresa tu fruta: ")
	fmt.Scanln(&opt)

	switch opt {
	case "Manzana":
		fmt.Println("La Manzana es buena")
	case "Pera":
		fmt.Println("La pera es mala")
	default:
		fmt.Printf("Ni se q es XD")
	}
}

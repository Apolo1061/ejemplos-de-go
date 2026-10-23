package main

import (
	"fmt"
)

func main() {
	var capitalde = map[string]string{
		"Chile":     "Santiago",
		"Venezuela": "Caracas",
	}

	var opt string
	fmt.Println("Capital de los paises")
	fmt.Print("Pais: ")
	fmt.Scanln(&opt)

	if opt == "Chile" {
		fmt.Println("Tienes permiso de acceder")
		fmt.Printf("La capital de %s es %s", opt, capitalde[opt])
	} else {
		println("\nNo tienes permiso de acceder")
	}
}

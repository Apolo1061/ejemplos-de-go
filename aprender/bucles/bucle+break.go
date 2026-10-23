package main

import "fmt"

func main() {

	var valor int = 0 // De manera mas facil se podria hacer valor := 0
	for {
		// A este if solo se entra si es que valor es igual o mayor a 10
		if valor >= 10 {
			break // break se usa para salir del bucle lo detienes
		}
		fmt.Println("valor = ", valor)
		valor++
	}

	fmt.Println("bucle terminado")
}

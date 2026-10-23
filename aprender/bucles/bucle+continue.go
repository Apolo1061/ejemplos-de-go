package main

import (
	"fmt"
)

func main() {
	valor := 0

	for valor < 15 {
		valor++ // Aumentamos valor
		// Por debajo es valor += 1

		if valor == 10 {
			continue
			// Al poner continue hace que el bucle deje de ejecutor lo que sigue
			// Lo que esto provoca es que el bucle cuando valor valga 10 nunca llegue
			// Al printf de abajo por lo tanto nunca se imprimira en stdout el valor
		}

		fmt.Printf("valor = %d \n", valor)
	}
}

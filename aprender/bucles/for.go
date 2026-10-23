package main

import (
	"fmt"
)

func main() {
	// Creacion de un bucle for clasic
	for i := 0; i < 5; i++ { // La alternativa a usar := seria var i int = 0
		fmt.Printf("Iteracion  %d \n", i) // Impresion del indicie del array
	}
}

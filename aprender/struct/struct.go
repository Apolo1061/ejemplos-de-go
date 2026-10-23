package main

import (
	"fmt"
)

type Humano struct { // Declaramos siempre el nombre de la struct en Mayuscula para decirle al compilador que es una struct accesible desde cualquier archivo
	Nombre   string
	Apellido string
	Pais     string
	Edad     int16
	Dinero   float32
}

func main() {
	var marco Humano // Le declaramos una variable local al struct
	/* Tambien se podria hacer marco := Humano
	// En ese caso se deja inicializada la struct
	// Todos los valores de las struct se inicializarian con los valores
	// Default de los tipos de datos
	// int = 0
	// string = ""
	// float32 = 0
	*/

	fmt.Print("[ System ] nombre> ")
	fmt.Scanln(&marco.Nombre)

	fmt.Print("[ System ] Apellido> ")
	fmt.Scanln(&marco.Apellido)

	fmt.Print("[ System ] Edad> ")
	fmt.Scanln(&marco.Edad)

	fmt.Print("[ System ] Dinero> ")
	fmt.Scanln(&marco.Dinero)

	fmt.Print("[ System ] Pais> ")
	fmt.Scanln(&marco.Pais)

	fmt.Println("\n\n\n\n\n")
	fmt.Println("Tus datos")
	fmt.Printf("%+v", marco)
}

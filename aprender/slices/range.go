package main

import (
	"fmt"
)

func main() {
	var cosas []string = []string{"Carro", "Reloj", "Celular"}

	for i, cosas := range cosas {
		fmt.Printf("cosas[ %d ] = %s \n", i, cosas)
	}
}

package main

import (
	"fmt"
)

func imprimir(nombres ...string) {
	fmt.Println("----- Impresion de varibale -----")
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}

	fmt.Println("----- Fin -----")
}

func main() {

	var nombres []string = []string{"Pedro", "Marco", "Sebas", "Antonio"}

	imprimir(nombres...)

}

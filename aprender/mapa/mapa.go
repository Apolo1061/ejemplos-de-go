package main

import (
	"fmt"
)

func main() {
	var capital = map[string]string{
		"Chile":     "Santiago",
		"Venezuela": "Caracas",
		"Basura":    "basura",
	}

	var opt string
	fmt.Println("Pais al que le quieres saber la capital")
	fmt.Print("Pais: ")
	fmt.Scanln(&opt)

	fmt.Println("La capital de chile es: ", capital[opt])

	delete(capital, "basura")
}

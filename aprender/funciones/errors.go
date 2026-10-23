package main

import (
	"errors"
	"fmt"
)

func dividir(a float64, b float64) (float64, error) {
	if b == 0 {
		fmt.Errorf("No se puede dividir por 0")
		return 0, errors.New("No se puede dividir por 0")
	}

	res := a / b
	return res, nil
}

func main() {
	var num1 float64 = 8
	var num2 float64 = 2

	cociente, error := dividir(num1, num2)
	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Printf("cociente = %2.f", cociente)
}

package main

import (
	"fmt"
)

func dividir(a float64, b float64) (float64, float64) {
	cociente := a / b
	resto := float64(int(a) % int(b))
	return cociente, resto
}

func main() {
	var num1 float64 = 4
	var num2 float64 = 2

	cociente, resto := dividir(num1, num2)

	fmt.Printf("cociente = %.2f \n", cociente)
	fmt.Printf("resto  = %.2f", resto)
}

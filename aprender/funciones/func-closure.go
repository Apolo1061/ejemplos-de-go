package main

import (
	"fmt"
)

func Crearcontador() func() int {
	var num int

	return func() int {
		num++
		return num
	}
}

func main() {
	contador := Crearcontador()
	a := contador()
	fmt.Println(a)

	b := contador()

	fmt.Printf("%d", b)

}

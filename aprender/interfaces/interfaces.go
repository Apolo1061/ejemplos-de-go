package main

import (
	"fmt"
)

type Perro struct{}
type Gato struct{}

type Hablador interface {
	hablar() string
}

func (p Perro) hablar() string {
	return "Guau"
}

func (g Gato) hablar() string {
	return "Miau"
}

func hacerHablar(h Hablador) {
	fmt.Println(h.hablar())
}

func main() {
	var rex Perro
	var kity Gato

	hacerHablar(rex)
	hacerHablar(kity)

}

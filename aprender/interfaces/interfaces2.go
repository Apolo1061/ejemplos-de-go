package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

func imprimirArea(f Forma) {
	fmt.Printf("Area: %.2f \n", f.Area())
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func main() {
	c := Circulo{Radio: 5.27}

	imprimirArea(c)
}

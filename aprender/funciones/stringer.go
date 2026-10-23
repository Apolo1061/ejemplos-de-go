package main

import (
	"fmt"
)

type Persona struct {
	Nombre   string
	Apellido string
	Edad     int
}

func (p Persona) String() string {
	return fmt.Sprintf("%s %s (%d años)", p.Nombre, p.Apellido, p.Edad)
}

func main() {
	persona1 := Persona{
		Nombre:   "Marco",
		Apellido: "Fernandez",
		Edad:     18,
	}

	fmt.Println(persona1)
}

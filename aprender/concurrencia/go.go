package main

import (
	"fmt"
	"time"
)

// Palabra clave go => GoRutine
// GoRutine -> Es cun hilo pero de ejecucion ligero y virtual
// channel 0> Es para comuncar a los GoRutines

func Hola(canal chan<- string) {
	time.Sleep(1 * time.Second)
	canal <- "Data a enviar"
	fmt.Println("Se envio el mensaje...")
}

func imprimirMensaje(canal <-chan string) {
	fmt.Println("Esperando mensaje...")
	msg := <-canal
	fmt.Println("Se recibio el mensaje")
	fmt.Println(msg)
}

func main() {
	canal := make(chan string)
	go Hola(canal)
	imprimirMensaje(canal)
}

package main

import (
	"fmt"
	"net"
)

type target struct {
	direccion string
	tiempo    int
}

func go_attack(t target) int {
	payload := []byte{0x00}
	conc, err := net.Dial("udp", t.direccion)
	if err != nil {
		fmt.Println("Error al crear socket")
		return 0
	}

	n, err := conc.Write(payload)
	if err != nil {
		fmt.Println("Error al enviar: ", err)
		return 0
	}

	conc.Close()
	return n
}

func main() {
	var tmp string
	var tmp2 string
	var tmp3 int

	fmt.Print("IP > ")
	fmt.Scanln(&tmp)
	if tmp == "" {
		fmt.Print("Ip vacia")
		return
	}

	fmt.Print("PORT > ")
	fmt.Scanln(&tmp2)
	if tmp2 == "" {
		fmt.Print("Puerto vacio")
		return
	}

	fmt.Print("TIME > ")
	fmt.Scanln(&tmp3)
	if tmp3 == 0 {
		fmt.Print("Tiempo vacio")
		return
	}

	tar := target{
		direccion: fmt.Sprintf("%s:%s", tmp, tmp2),
		tiempo:    tmp3,
	}

	n := go_attack(tar)
	if n <= 0 {
		fmt.Print("Error de envio de paquetes")
		fmt.Printf("\nSe enviaron %d paquetes", n)
		return
	}

	println("Se mandaron ", n, "paquetes")
}

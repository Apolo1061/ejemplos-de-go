package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "Hola"
	string2 := " Mundo!"

	concat := string1 + string2
	mayus := strings.ToUpper(concat)

	fmt.Printf("%s", concat)
	fmt.Printf("\n %s", mayus)
}

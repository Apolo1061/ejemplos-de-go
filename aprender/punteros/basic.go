package main

import (
	"fmt"
)

func incrementar(num *int) {
	*num++
}

func main() {
	x := 0
	fmt.Println(x)

	incrementar(&x)
	fmt.Println(x)
}

package main

import (
	"fmt"
)

func main() {
	var slice []int = []int{1, 2, 3}

	fmt.Println(slice[0:])

	slice = append(slice, 4)

	fmt.Println(slice[0:])
}

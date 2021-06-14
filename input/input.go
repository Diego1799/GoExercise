package main

import (
	"fmt"
)

func main() {
	var edad int
	fmt.Println("Ingresa edad")
	fmt.Scanln(&edad)
	fmt.Println("Su edad es:", edad)
}

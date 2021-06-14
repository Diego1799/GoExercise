package main

import (
	"fmt"
	"strconv"
)

func main() {
	nombre := "Diego"
	edad := 22
	fmt.Println("Mi nombre es " + nombre + " y tengo " + strconv.Itoa(edad) + " años")
}

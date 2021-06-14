package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	nombre_lento("Diego")
}

func nombre_lento(name string) {
	letras := strings.Split(name, "")

	for _, letra := range letras {
		time.Sleep(1 * time.Second)
		fmt.Println(letra)
	}
}

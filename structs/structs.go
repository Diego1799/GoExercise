package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func main() {
	Diego := User{edad: 21, nombre: "Diego", apellido: "Ramirez"}
	fmt.Println(Diego)
}

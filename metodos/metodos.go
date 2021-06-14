package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (usuario User) nombre_completo() string {
	return usuario.nombre + " " + usuario.apellido
}

func (usuario *User) set_name(name string) {
	usuario.nombre = name
}

func main() {
	Diego := User{edad: 21, nombre: "Diego", apellido: "Ramirez"}

	Diego.set_name("El Cacha")
	fmt.Println(Diego.nombre_completo())

}

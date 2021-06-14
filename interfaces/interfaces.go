package main

import "fmt"

type User interface {
	Permisos() int
	Nombre() string
}

type Admin struct {
	name string
}

func (this Admin) Permisos() int {
	return 5
}

func (this Admin) Nombre() string {
	return this.name
}

type Cliente struct {
	name string
}

func (this Cliente) Permisos() int {
	return 3
}

func (this Cliente) Nombre() string {
	return this.name
}

func auth(user User) string {
	if user.Permisos() > 4 {
		return user.Nombre() + " Tiene permisos de admin"
	} else {
		return user.Nombre() + " No es admin"
	}
}

func main() {
	admin := Admin{"Diego"}
	cliente := Cliente{"Cachaco"}
	fmt.Println(auth(admin))
	fmt.Println(auth(cliente))
}

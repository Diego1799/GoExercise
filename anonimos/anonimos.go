package main

import "fmt"

type Human struct {
	nombre string
}

type Tutor struct {
	Human
}

func main() {
	tutor := Tutor{Human{"Diego"}}

	fmt.Println(tutor.nombre)
}

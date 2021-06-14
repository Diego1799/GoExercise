package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("./example.log")
	if err != nil {
		fmt.Println("Hubo un error leyendo")
	} else {
		fmt.Println(string(file))
	}
}

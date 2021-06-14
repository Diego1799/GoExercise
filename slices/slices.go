package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)

	arreglo := [3]int{1, 2, 3}
	slice2 := arreglo[0:]
	fmt.Println(slice2)
}

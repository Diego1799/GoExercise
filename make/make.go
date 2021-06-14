package main

import "fmt"

func main() {
	slice := make([]int, 0, 3)
	slice = append(slice, 23)
	slice = append(slice, 24)
	slice = append(slice, 25)
	fmt.Println(slice)
	fmt.Println(cap(slice))

}

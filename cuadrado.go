package main

import "fmt"

func main() {
	var base float64
	fmt.Print("Longitud de la base: ")
	fmt.Scan(&base)

	area := base * base
	fmt.Println("Área del cuadrado: ", area)
}

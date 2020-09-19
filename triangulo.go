package main

import "fmt"

func main() {
	var base float64
	var altura float64

	fmt.Print("Base del triángulo: ")
	fmt.Scan(&base)
	fmt.Print("Altura del triángulo: ")
	fmt.Scan(&altura)

	area := (base * altura) / 2

	fmt.Print("Área del triángulo: ", area)
}

package main

import "fmt"

func main() {
	var baseMenor float64
	var baseMayor float64
	var altura float64

	fmt.Print("Base menor del trapecio: ")
	fmt.Scan(&baseMenor)
	fmt.Print("Base mayor del trapecio: ")
	fmt.Scan(&baseMayor)
	fmt.Print("Altura del trapecio: ")
	fmt.Scan(&altura)

	area := (baseMayor + baseMenor) / 2 * altura

	fmt.Print("Área del trapecio: ", area)
}

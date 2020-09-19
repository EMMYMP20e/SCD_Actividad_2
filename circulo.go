package main

import "fmt"

func main() {
	var radio float64

	fmt.Print("Radio del círculo: ")
	fmt.Scan(&radio)

	area := 3.14 * (radio * radio)

	fmt.Println("Área del círculo: ", area)
}

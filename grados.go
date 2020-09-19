package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Print("Grados Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celcius := ((fahrenheit - 32) * 5 / 9)
	fmt.Println("Grados Celcius: ", celcius)
}

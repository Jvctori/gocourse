package main

import "fmt"

func main() {
	// Variables declaration
	// SOMA!
	var a, b int = 10, 3
	var result int
	result = a + b
	fmt.Println("Adittion:", result)

	// SUBTRAÇÃO
	result = a - b
	fmt.Println("Subtraction:", result)

	// DIVISÃO

	result = a / b
	fmt.Println("Division:", result)

	// Multplicação
	result = a * b
	fmt.Println("Multiplication:", result)

	// Modulo
	result = a % b
	fmt.Println("Module:", result)

	var p float64 = 22 / 7
	fmt.Println("O interpretador do go não irá retornar a divisão exata (em float) pois os 2 valores, apesar de terem sidos delcarados como float, não carregam .0", p)
	fmt.Println(p)
	var q float64 = 22 / 7.0
	fmt.Println("o 7 foi declarado como: 7.0, a divisão será precisa:", q)

}

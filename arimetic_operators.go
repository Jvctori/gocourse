package main

import (
	"fmt"
	"math"
)

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
	fmt.Println("O interpretador do go não irá retornar a divisão exata (em float) pois os 2 valores, apesar de terem sidos delcarados como float, não carregam: .0, p =", p)

	var q float64 = 22 / 7.0
	fmt.Println("o 7 foi declarado como: 7.0, a divisão será precisa: q =", q)
	// É uma boa praticar colocar as operações aritmetricas entre ()
	// exemplo: (1 + 1)
	// Cuidado que o tipo de dado suporta uma quantidade limitada de valor
	// se voce passar um valor maior do que o suportado em float32...
	// ...ocorrerá um erro: se configura: OVERFLOW

	// OVERFLOW CONCEITO:
	/*var maxInt int64 = 92233221314127231823912746123617
	fmt.Println(maxInt)*/
	// output: ./arimetic_operators.go:41:21: cannot use 92233221314127231823912746123617 (untyped int constant) as int64 value in variable declaration (overflows)
	// o valor maximo para int64 é: 9223372036854775807
	// OVERFLOW COM INT64
	var maxInt int64 = 9223372036854775807 // int64 max value
	fmt.Println("Valor maximo do int64:", maxInt)
	maxInt = (maxInt + 1)
	fmt.Println("O que acontece se adicionar +1 ao int64:", maxInt)
	// OVERFLOW com unsigned int (inteiros não negativos)
	var uMaxInt uint64 = 18446744073709551615 // uint64 max value
	fmt.Println("Valor maximo de unsigned int (uint64):", uMaxInt)
	uMaxInt = (uMaxInt + 1)
	fmt.Println("O que acontece se adicionar +1 ao uint64:", uMaxInt)

	var smallFloat float64 = 1.0e-323
	fmt.Println(smallFloat)
	smallFloat = (smallFloat / math.MaxFloat64)
	fmt.Println(smallFloat)
	// UNDERFLOW with floating numbers:
	// output foi 0 pois é uma divisão muito grande, ou seja:...
	// não há possiblidade do float64 representar esse valor muito pequeno
}

package basics

import "fmt"

// se quisermos declarar uma variavel globalmente, use essa forma:
// var variavel string = ""
// var numero int = 10
// ...
// ou var middleName = "String", deixando o compilador tipar
var middleName string = "Silva"

func test2() {
	var age int
	var name string = "John"
	var name1 = "Jane"
	count := 10
	lastName := "Victor"

	// Default values:
	// Numeric Types: 0
	// Boolean Types: false
	// String Type: ""
	// Pointers, slices, maps, functions, and structs: nil

	fmt.Println(name)
	fmt.Println(name1)
	fmt.Println(count)
	fmt.Println(lastName)
	fmt.Scan(&age)
	printname()
	fmt.Println(middleName)

}

func printname() {
	firstName := "Michael"
	fmt.Println(firstName)
}

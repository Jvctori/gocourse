package basics

import "fmt"

// Exemplo de name convention struct:
type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func test3() {
	// PascalCase:
	// exmp: CalculateArea, UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case:
	// exmp: user_id, first_name, http_requests
	// Files, directory...

	// UPPERCASE
	// use in constants
	// exmp: PI

	// mixedCase
	// pode ser usado em variaveis
	// could be used in variables
	// exmp: javaScript, htmlDocument, isValid

	// package names: mantenha em lower case, sem underscore "_"
	// obs: Padronize! Se for usar snake_case para arquivos, use mixedCase para variaveis
	// Exemplo de name convention constante:
	const MAXRETRIES = 5
	// Exemplo de name convention variavel:
	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)
	fmt.Println(MAXRETRIES)
}

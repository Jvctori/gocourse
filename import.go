package main

import (
	"fmt"
	"net/http"
)

// Three shaking and its application
// three shaking remove modulos de bibliotecas que não estão sendo usados
// go compiler j́á faz isso para você
func main() {
	fmt.Println("Hello, Go Standard Library")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	defer resp.Body.Close()

	fmt.Println("Http Response Status: ", resp.Status)
}

package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Go Standard Library")

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	defer resp.Body.Close()

	fmt.Println("Http Response Status: ", resp.Status)
}

package main

//go run main.go
//go build

import (
	"fmt"
)

func main() {
	n, error := fmt.Println("Hello World")
	if error != nil {
		fmt.Println("Salida de n: ", n)
		fmt.Println("Salida de error: ", error)
	}
}

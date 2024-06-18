package main

import "fmt"

func main() {
	// Chamando a função original
	fmt.Println("Hello, World!")
	
	// Chamando a função sayHello do main2.go
	sayHello()
	
	// Chamando a função DemonstrarFundamentos do fundamentos.go
	DemonstrarFundamentos()

	// Chamando a função sayHello do main3.go
	funcao()
}

package main

import "fmt"

// Função principal
func main() {
	// Variáveis
	var nome string = "João"
	var idade int = 30
	var altura float64 = 1.75
	var ativo bool = true

	// Imprimindo variáveis
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Ativo:", ativo)

	// Estruturas de controle: if-else
	if idade >= 18 {
		fmt.Println(nome, "é maior de idade.")
	} else {
		fmt.Println(nome, "é menor de idade.")
	}

	// Laço de repetição: for
	fmt.Println("Contagem de 1 a 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Função
	resultado := soma(5, 7)
	fmt.Println("A soma de 5 e 7 é:", resultado)

	// Array
	numeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Números no array:")
	for _, numero := range numeros {
		fmt.Println(numero)
	}
}

// Função para somar dois números
func soma(a int, b int) int {
	return a + b
}

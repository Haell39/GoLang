package main

import "fmt"

// Variáveis e Tipos de Dados
func demoVariaveisETipos() {
	var numero int = 42
	var texto string = "Olá, Go!"
	var booleano bool = true

	fmt.Println("Número:", numero)
	fmt.Println("Texto:", texto)
	fmt.Println("Booleano:", booleano)
}

// Estruturas Condicionais
func demoCondicionais() {
	numero := 10
	if numero > 5 {
		fmt.Println("O número é maior que 5")
	} else {
		fmt.Println("O número é 5 ou menor")
	}
}

// Laços
func demoLacos() {
	for i := 0; i < 5; i++ {
		fmt.Println("Contagem:", i)
	}

	numero := 0
	for numero < 3 {
		fmt.Println("Número:", numero)
		numero++
	}
}

// Funções
func somar(a int, b int) int {
	return a + b
}

func demoFuncoes() {
	resultado := somar(3, 4)
	fmt.Println("Resultado da soma:", resultado)
}

// Arrays e Slices
func demoArraysESlices() {
	// Arrays
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Slices
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice:", slice)
}

// Mapas
func demoMapas() {
	idade := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	fmt.Println("Idades:", idade)
	fmt.Println("Idade de Alice:", idade["Alice"])
}

// Structs
type Pessoa struct {
	Nome  string
	Idade int
}

func demoStructs() {
	pessoa := Pessoa{"Carlos", 28}
	fmt.Println("Pessoa:", pessoa)
	fmt.Println("Nome:", pessoa.Nome)
	fmt.Println("Idade:", pessoa.Idade)
}

// Função principal para demonstrar todos os fundamentos
func DemonstrarFundamentos() {
	fmt.Println("=== Variáveis e Tipos de Dados ===")
	demoVariaveisETipos()

	fmt.Println("\n=== Estruturas Condicionais ===")
	demoCondicionais()

	fmt.Println("\n=== Laços ===")
	demoLacos()

	fmt.Println("\n=== Funções ===")
	demoFuncoes()

	fmt.Println("\n=== Arrays e Slices ===")
	demoArraysESlices()

	fmt.Println("\n=== Mapas ===")
	demoMapas()

	fmt.Println("\n=== Structs ===")
	demoStructs()
}

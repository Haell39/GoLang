package main

import (
	"fmt"
	"errors"
	"math"
	"os"
)

// Função principal, onde a execução do programa começa
func Fundamentosback() {
	// Declaração de variáveis e tipos básicos
	var myInt int = 10
	var myFloat float64 = 3.14
	myString := "Hello, Go!" // Inferência de tipo
	myBool := true

	fmt.Println("Integer:", myInt)
	fmt.Println("Float:", myFloat)
	fmt.Println("String:", myString)
	fmt.Println("Boolean:", myBool)

	// Operadores aritméticos
	sum := myInt + 2
	difference := myFloat - 1.14
	product := myInt * 2
	quotient := myFloat / 2
	remainder := myInt % 3

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	// Operadores relacionais e lógicos
	isEqual := myInt == 10
	isNotEqual := myFloat != 3.14
	isGreater := myInt > 5
	isLess := myFloat < 5.0
	isTrueAnd := myBool && (myInt < 20)
	isTrueOr := myBool || (myInt > 20)

	fmt.Println("Is Equal:", isEqual)
	fmt.Println("Is Not Equal:", isNotEqual)
	fmt.Println("Is Greater:", isGreater)
	fmt.Println("Is Less:", isLess)
	fmt.Println("Is True And:", isTrueAnd)
	fmt.Println("Is True Or:", isTrueOr)

	// Controle de fluxo: if, else, switch
	if myInt > 5 {
		fmt.Println("myInt is greater than 5")
	} else {
		fmt.Println("myInt is not greater than 5")
	}

	// Switch case em Go
	switch myInt {
	case 10:
		fmt.Println("myInt is 10")
	case 20:
		fmt.Println("myInt is 20")
	default:
		fmt.Println("myInt is neither 10 nor 20")
	}

	// Estruturas de repetição: for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Loop index:", i)
	}

	// Arrays e slices
	myArray := [3]int{1, 2, 3}
	mySlice := []int{4, 5, 6}

	fmt.Println("Array:", myArray)
	fmt.Println("Slice:", mySlice)

	// Adicionando elementos a um slice
	mySlice = append(mySlice, 7, 8)
	fmt.Println("Updated Slice:", mySlice)

	// Iteração sobre slice
	for index, value := range mySlice {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}

	// Funções: definição, argumentos e retorno de valores
	result := add(5, 3)
	fmt.Println("Add result:", result)

	// Funções com retorno múltiplo
	quot, rem := divide(10, 3)
	fmt.Printf("Quotient: %d, Remainder: %d\n", quot, rem)

	// Structs: definição e métodos
	person := Person{name: "John", age: 30}
	person.greet()

	// Manipulação de ponteiros
	increment(&myInt)
	fmt.Println("Incremented myInt:", myInt)

	// Tratamento de erros
	if err := checkAge(person.age); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Age is valid")
	}

	// Entrada e saída: leitura do terminal
	var input string
	fmt.Print("Enter some text: ")
	fmt.Scanln(&input)
	fmt.Println("You entered:", input)

	// Manipulação de arquivos
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, file handling in Go!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println("File written successfully")
}

// Função simples para somar dois inteiros
func add(a int, b int) int {
	return a + b
}

// Função para dividir dois inteiros com retorno de quociente e resto
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

// Struct em Go
type Person struct {
	name string
	age  int
}

// Método para a struct Person
func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

// Função que usa ponteiros para incrementar um valor
func increment(x *int) {
	*x = *x + 1
}

// Função que retorna um erro se a idade for inválida
func checkAge(age int) error {
	if age < 18 {
		return errors.New("age must be 18 or older")
	}
	return nil
}

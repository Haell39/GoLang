package main

import (
	"fmt"
	"os"
	"io"
)

// Declaração de variáveis e tipos de dados
var (
	// Tipos básicos
	myInt     int     = 10
	myFloat   float64 = 3.14
	myString  string  = "Hello, Go!"
	myBool    bool    = true
)

func main() {
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
	isEqual := (myInt == 10)
	isNotEqual := (myFloat != 3.14)
	isGreater := (myInt > 5)
	isLess := (myFloat < 5.0)
	isTrueAnd := (myBool && (myInt < 20))
	isTrueOr := (myBool || (myInt > 20))

	fmt.Println("Is Equal:", isEqual)
	fmt.Println("Is Not Equal:", isNotEqual)
	fmt.Println("Is Greater:", isGreater)
	fmt.Println("Is Less:", isLess)
	fmt.Println("Is True And:", isTrueAnd)
	fmt.Println("Is True Or:", isTrueOr)

	// Estruturas de controle de fluxo (if, else, switch)
	if myInt > 5 {
		fmt.Println("myInt is greater than 5")
	} else {
		fmt.Println("myInt is not greater than 5")
	}

	switch myInt {
	case 10:
		fmt.Println("myInt is 10")
	case 20:
		fmt.Println("myInt is 20")
	default:
		fmt.Println("myInt is neither 10 nor 20")
	}

	// Estruturas de repetição (for, range)
	for i := 0; i < 5; i++ {
		fmt.Println("Loop index:", i)
	}

	// Arrays e slices
	myArray := [3]int{1, 2, 3}
	mySlice := []int{4, 5, 6, 7}

	fmt.Println("Array:", myArray)
	fmt.Println("Slice:", mySlice)

	// Iteração sobre slice com range
	for index, value := range mySlice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Funções, incluindo passagem de argumentos e retorno de valores
	result := add(5, 3)
	fmt.Println("Add result:", result)

	// Structs e métodos
	p := Person{name: "John", age: 30}
	fmt.Println("Person:", p)
	p.greet()

	// Ponteiros e manipulação de memória
	x := 10
	fmt.Println("Original x:", x)
	increment(&x)
	fmt.Println("Incremented x:", x)

	// Entrada e saída
	var userInput string
	fmt.Println("Enter some text:")
	fmt.Scanln(&userInput)
	fmt.Println("You entered:", userInput)

	// Manipulação de arquivos com os pacotes os e io
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.WriteString(file, "Hello, file handling in Go!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("File written successfully")
}

// Função que soma dois inteiros
func add(a int, b int) int {
	return a + b
}

// Definição de struct
type Person struct {
	name string
	age  int
}

// Método associado à struct Person
func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

// Função que incrementa o valor de um inteiro usando um ponteiro
func increment(x *int) {
	*x = *x + 1
}

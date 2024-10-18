package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int16 = 32767 // there also int and int64(for bigger numbers)
	// int16 has the maximum capacity of 32767, bigger than that cause error
	var intNum2 int8 = 127 // int8: (-128, 127)
	var intNum3 uint8 = 0  // unint8: (0, 255) *unint is only for positives
	fmt.Println(intNum)
	fmt.Println(intNum2)
	fmt.Println(intNum3)

	var floatNum float64 = 12345678.9  //  precisão de até 15-16
	var floatNum2 float32 = 12345678.9 //float32 precisão de cerca de 7 dígitos
	fmt.Printf("%.2f\n", floatNum)
	fmt.Printf("%.2f\n", floatNum2)

	//Is not possible do operations with different types

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// var result float32 = floatNum32 + intNum32 WILL result ERROR!!
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// Integer division will result an integer number

	var intNum1 int = 3
	var intNum4 int = 2
	fmt.Println(intNum1 / intNum4) // result will be 1 because go will round
	fmt.Println(intNum1 % intNum4) // result will also be 1, but here because % means the rest

	//String

	var myString string = "Hello, World!"
	var myString2 string = "Hello," + " " + "World!"
	var myString3 string = `Backwards quotes 
let you format your text without using \n notation`

	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(myString3)

	var myString5 string = "Lótus"

	fmt.Println(len(myString5))                    // returns 6 (wrong)
	fmt.Println(utf8.RuneCountInString(myString5)) // returns 5 (Right)

	/*

			package main

			import (
			    "fmt"
			    "unicode/utf8"
			)

			func main() {
			    //String com caracteres do alfabeto inglês e um acento
			    str := "Café"

			    // Contagem usando len
			    fmt.Println("Número de bytes (len):", len(str))
			// Saída: 5

			    // Contagem usando utf8.RuneCountInString
			    fmt.Println("Número de runes (utf8.RuneCountInString):", utf8.RuneCountInString(str))
			// Saída: 4
			}

		OBS: Ou seja Para strings que contêm apenas caracteres do alfabeto inglês sem acentuação o *len* serve perfeitamente mas para caracteres acentuados, emojis ou caracteres de outros idiomas é melhor usar *utf8.RuneCountInString*

	*/

	var myRune rune = 'a'
	fmt.Println(myRune) // not a typo, very strange see later

	var myBoolean bool = true
	fmt.Println(myBoolean)

	// we can omite the typo if we do this:

	var myVar = "text"
	fmt.Println(myVar)
	// or we can do this:

	myVar2 := "More, Text"
	fmt.Println(myVar2)

	//we can initialize more than one variable:

	var1, var2, var3 := 1, 2, 3
	fmt.Println(var1, var2, var3)

	// When is not obvious its better to specify the type of the var! For Exp:
	/*
		var myVar string = lee() //is not obvious what the function lee is, so is better specify it!
		fmt.Println(myVar)
	*/

	const myCounst string = "const Value!" //everything we learn aply to const except we can not keep modifying the const once assigned
	fmt.Println(myCounst)

	const pi float64 = 3.1415926 // there s no need to change this

}

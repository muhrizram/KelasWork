package main

import (
	"errors"
	"fmt"
)

func main() {
	var first string = "This is First thing."
	second := "This is The Second."

	fmt.Println(first)
	fmt.Println(second)

	// Primitive Data Type in GO Lang is Boolean, String, Int, Float

	// Boolean Example
	thisBool := true

	fmt.Printf("Type: %T Value: %v\n", thisBool, thisBool)

	// Integer Example
	intEx := int(5)
	intEx1 := int32(5)
	intEx2 := int64(5)

	fmt.Printf("Type: %T Value: %v\n", intEx, intEx)
	fmt.Printf("Type: %T Value: %v\n", intEx1, intEx1)
	fmt.Printf("Type: %T Value: %v\n", intEx2, intEx2)

	// Float Example
	floatEx := float32(6)
	floatEx1 := float64(6)

	fmt.Printf("Type: %T Value: %v\n", floatEx, floatEx)
	fmt.Printf("Type: %T Value: %v\n", floatEx1, floatEx1)

	// Bytes Example
	bytesVar := byte(255)
	bytesVar2 := []byte("Trying to be String")

	fmt.Printf("Type: %T Value: %v\n", bytesVar, bytesVar)
	fmt.Printf("Type: %T Value: %v\n", bytesVar2, bytesVar2)

	// Rune Example
	runeEx := '😁'

	fmt.Printf("Type: %T Value: %v\n", runeEx, runeEx)

	// Complex / Imaginer Value Example
	complexEx := -7 - 1i

	fmt.Printf("Type: %T Value: %v\n", complexEx, complexEx)

	// Error Example
	errVar := errors.New("Error Detected")

	fmt.Printf("Type: %T Value: %v\n", errVar, errVar)

	// Interface Example
	var myInterfaceVar interface{}
	var myInterfaceVar1 interface{}

	myInterfaceVar = 2
	fmt.Printf("Type: %T Value: %v\n", myInterfaceVar, myInterfaceVar)
	myInterfaceVar1 = true
	fmt.Printf("Type: %T Value: %v\n", myInterfaceVar1, myInterfaceVar1)
}

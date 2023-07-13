package main

import (
	"errors"
	"fmt"
)

func main() {

	// Numbers
	var number1 int8 = 127
	fmt.Println(number1)
	
	var number2 int16 = 10000
	fmt.Println(number2)

	var number3 uint32 = 10000
	fmt.Println(number3)

	// int32 = rune
	var number4 rune = 123456
	fmt.Println(number4)

	// byte = uint8
	var number5 byte = 127
	fmt.Println(number5)

	var number6 float32 = 231323.23
	fmt.Println(number6)

	var number7 float64 = 10201020000000000000.45
	fmt.Println(number7)

	number8 := 12313.49
	fmt.Println(number8)

	// Strings
	var str string = "Text"
	fmt.Println(str)

	str2 := "Text"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// Zero data-types

	// equal to ""
	var text string
	fmt.Println(text)

	// equal to 0
	var number int16
	fmt.Println(number)

	// equal to false
	var boolean bool
	fmt.Println(boolean)

	// equal to nil
	var error1 error
	fmt.Println(error1)

	// Error type
	var error2 error = errors.New("Internal error")
	fmt.Println(error2)
}
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create an int array with 5 positions
	var array1[5] int
	fmt.Println(array1)

	// Create a string array with 5 positions
	var array2[5] string
	fmt.Println(array2)

	// Define array of string and setting the values
	array3 := [5]string{"Pos1", "Pos2", "Pos3", "Pos4", "Pos5"}
	fmt.Println(array3)

	// This array bellow is not dynamic
	array4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array4)

	// This is not an array, but points to one
	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	// They are not the same type
	fmt.Println(reflect.TypeOf(array4))
	fmt.Println(reflect.TypeOf(slice))

	// Append elements to the end of the array
	slice = append(slice, 18)
	fmt.Println(slice)

	// Slice the array, the last position is not included
	slice2 := array3[1:3]
	fmt.Println(slice2)

	// The slice2 points to array3
	array3[1] = "New Pos"
	fmt.Println(slice2)

	// Internals arrays
	// Create an slice with types float32, 10 positions, and 11 max positions
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	fmt.Println(len(slice3)) // Get length of slice
	fmt.Println(cap(slice3)) // Get max length of slice

	slice3 = append(slice3, 1)
	fmt.Println(slice3)

	// If the array capacity is exceeded, golang creates a new array with a doubled capacity
	slice3 = append(slice3, 1)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	// If the last parameter is not defined, golang set the capacity to number of elements
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
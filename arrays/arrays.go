package main

import (
	"fmt"
)

/*
 * An array is a numbered sequence of elements of a specific length.
 * In Go, an array is a numbered sequence of elements of a specific length.
 * Arrays are indexed starting from 0.
 * The type of elements and length are both part of the array's type.
 * By default, an array is zero-valued, which for ints means 0s.
 * Arrays have a fixed size.
 * Arrays are useful when planning the detailed layout of memory.
 */
func main() {
	var a [5]int // an array of 5 integers which have a default value of 0

	fmt.Println("start with: ", a)

	a[4] = 100 // set the value at index 4 to 100

	fmt.Println("set: ", a)

	fmt.Println("the length of the array is", len(a))

	// An array can be initialized in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("this was initialized in one line: ", b)

	// Arrays are one-dimensional, but you can compose types to build multi-dimensional data structures
	var twoDimensionalArray [3][3]int

	for i := 0; i < len(twoDimensionalArray); i++ {
		for j := 0; j < len(twoDimensionalArray[i]); j++ {
			twoDimensionalArray[i][j] = i + j
		}
	}

	fmt.Println("two dimensional array: ", twoDimensionalArray)
}

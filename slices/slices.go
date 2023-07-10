package main

import (
	"fmt"
)

func main() {
	var s []string

	s = make([]string, 3)

	fmt.Println("initial slice", s)
	fmt.Println("initial length", len(s))
	fmt.Println("initial capacity", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("filled slice:", s)

	fmt.Println("new length", len(s))
	fmt.Println("new capacity", cap(s))

	s = append(s, "d")

	fmt.Println("appended slice:", s)
	fmt.Println("appended length", len(s))
	fmt.Println("appended capacity", cap(s))

	s = append(s, "e", "f")

	fmt.Println("appended multiple:", s)
	fmt.Println("appended multiple length", len(s))
	fmt.Println("appended multiple capacity", cap(s))

	c := make([]string, len(s))

	fmt.Println("new slice:", c)
	fmt.Println("new length", len(c))
	fmt.Println("new capacity", cap(c))

	copy(c, s)

	fmt.Println("copied slice:", c)
	fmt.Println("copied length", len(c))
	fmt.Println("copied capacity", cap(c))

	l := c[2:5]

	fmt.Println("sliced slice:", l)
	fmt.Println("sliced length", len(l))
	fmt.Println("sliced capacity", cap(l))

	l = c[:5]

	fmt.Println("sliced slice:", l)
	fmt.Println("sliced length", len(l))
	fmt.Println("sliced capacity", cap(l))

	l = c[2:]

	fmt.Println("sliced slice:", l)
	fmt.Println("sliced length", len(l))
	fmt.Println("sliced capacity", cap(l))

	t := []string{"g", "h", "i"}

	fmt.Println("declared slice:", t)
	fmt.Println("declared length", len(t))
	fmt.Println("declared capacity", cap(t))

	twoDimensionalSlice := make([][]int, 3)

	for i := 0; i < len(twoDimensionalSlice); i++ {
		innerSliceLength := i + 1
		twoDimensionalSlice[i] = make([]int, innerSliceLength)
		for j := 0; j < innerSliceLength; j++ {
			twoDimensionalSlice[i][j] = i + j
		}
	}

	fmt.Println("two dimensional slice:", twoDimensionalSlice)
	fmt.Println("two dimensional length", len(twoDimensionalSlice))
	fmt.Println("two dimensional capacity", cap(twoDimensionalSlice))
}

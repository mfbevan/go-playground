package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")

	_time := time.Now()

	fmt.Println("The time is", _time)
	fmt.Println("The day is", _time.Weekday())

	switch _time.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	if _time.Hour() < 12 {
		fmt.Println("It's before noon")
	} else {
		fmt.Println("It's after noon")
	}
}

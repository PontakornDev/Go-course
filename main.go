package main

import "fmt"

var numberInt1, numberInt2 int = 0, 1

func main() {
	numberFoalt := 15.5
	msg := "Hello"
	fmt.Println(numberInt1)
	fmt.Println(numberInt2)
	fmt.Println(numberFoalt)
	fmt.Println(float64(numberInt1) + numberFoalt)
	fmt.Println(msg + " Pontakorn")
	fmt.Println("My money = ", float64(numberInt2)+numberFoalt)
}

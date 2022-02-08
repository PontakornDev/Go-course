package main

import "fmt"

var productName [4]string
var price [4]float32

func main() {
	productName[0] = "A"
	productName[1] = "B"
	productName[2] = "C"
	productName[3] = "D"

	price := [4]float32{1, 2, 3, 4}
	fmt.Println(productName)
	fmt.Println(price)
}

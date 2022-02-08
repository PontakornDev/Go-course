package main

import "fmt"

func hello() {
	fmt.Println("Hello Kuy")
}

func plus(value1 int, value2 int) {
	result := value1 + value2
	fmt.Println("result1 = ", result)
}

func plus2(value1 int, value2 int) int {
	return value1 + value2
}

func plus3(value1, value2, value3 int) int {
	return value1 + value2 + value3
}
func main() {
	hello()
	plus(1, 2)
	result2 := plus2(3, 4)
	fmt.Println("result2 = ", result2)
	result3 := plus3(5, 6, 7)
	fmt.Println("result3 = ", result3)
}

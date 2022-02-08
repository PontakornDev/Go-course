package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		input = strings.TrimRight(input, "\r\n")
		message, err := fmt.Printf("'%s' ERR: not string must number only \n", input)
		if err != nil {
			fmt.Println(err)
		}
		panic(message)
	}
	return value
}

func getOperator() string {
	fmt.Print("operator is (+ - * /): ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func delete(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func main() {
	value1 := getInput(" Value1 = ")
	value2 := getInput(" Value2 = ")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = delete(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		fmt.Println("Incorrect Operator")
	}
	fmt.Printf("result is %v", result)
}

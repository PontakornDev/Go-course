package main

import "fmt"

var nameColor string

func main() {
	fmt.Println("Input Name Color")
	fmt.Scanf("%s", &nameColor)
	switch nameColor {
	case "blue":
		fmt.Println("#0000FF")
	case "green":
		fmt.Println("#008000")
	case "pink":
		fmt.Println("#FFC0CB")
	case "yellow":
		fmt.Println("#FFFF00")
	default:
		fmt.Println("no color name")
	}
}

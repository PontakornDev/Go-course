package main

import "fmt"

var point int

func calGrade(point int) {
	if point >= 80 && point <= 100 {
		fmt.Println("A")
	} else if point >= 70 && point < 80 {
		fmt.Println("B")
	} else if point >= 60 && point < 70 {
		fmt.Println("C")
	} else if point >= 50 && point < 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}

func main() {
	fmt.Println("Grade calculator")
	fmt.Scanf("%d", &point)
	fmt.Println("Point = ", point)
	calGrade(point)

}

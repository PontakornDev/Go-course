package main

import "fmt"

var courseName []string

func main() {
	courseName = []string{"Java", "Python"}
	fmt.Println(courseName)
	courseName = append(courseName, "C", "C#", "Javascript")
	fmt.Println(courseName)

	courseWeb := courseName[3:5]
	fmt.Println(courseWeb)
	courseWeb = courseName[:3]
	fmt.Println(courseWeb)
}

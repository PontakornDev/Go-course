package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product = ", product)

	//! add
	product["A"] = 1
	product["B"] = 2
	product["C"] = 3
	fmt.Println("add = ", product)

	//!delete
	delete(product, "A")
	fmt.Println("delete = ", product)

	//!update
	product["C"] = 5555
	fmt.Println("update = ", product)

	values1 := product["C"]
	fmt.Println("values1 = ", values1)

	courseName := map[string]string{"101": "Java", "102": "Python"}
	fmt.Println(courseName)
}

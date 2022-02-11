package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {
	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "Pontakorn",
		phone:        "0620409695",
	}
	employee2 := employee{
		employeeID:   "102",
		employeeName: "Par",
		phone:        "0801334016",
	}
	employee3 := employee{
		employeeID:   "103",
		employeeName: "Ford",
		phone:        "0876562007",
	}

	employeeList = append(employeeList, employee1, employee2, employee3)
	fmt.Println("employee = ", employeeList)
}

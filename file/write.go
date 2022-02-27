package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data1 := []byte("Hi\n Pontakorn")
	err := os.WriteFile("data.txt", data1, 0644)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("EmployeeName")
	check(err)

	defer f.Close()

	data2 := []byte("Pontakorn\n Changpat")
	os.WriteFile("EmployeeName.txt", data2, 0644)
}

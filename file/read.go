package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	flie, err := os.Open("testData.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(flie)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, (","))
		fmt.Println(item[0])
		// fmt.Printf("Line %d : %s\n", count, line)

		count++
	}
}

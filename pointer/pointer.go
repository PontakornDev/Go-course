package main

import "fmt"

func zeroValue(iValue int) {
	iValue = 0
}

func zeroPointer(iPointer *int) {
	*iPointer = 0
}

func main() {
	i := 1
	fmt.Println("i = ", i)
	zeroValue(i)
	fmt.Println("i value form zero function zerovalue = ", i)
	zeroPointer(&i)
	fmt.Println("i address form zero function zeroPointer = ", &i)

}

package main

import (
	"fmt"
	"time"
)

func process1(ch chan string, data string) {
	ch <- data
}

func main() {
	ch := make(chan string)
	go process1(ch, "Data1")
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}

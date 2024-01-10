package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		fmt.Println("hello world")
	}()

	time.Sleep(time.Millisecond * 100)
}

func returnVal() string {
	return "hello world"
}

package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	str := strconv.Itoa(10)

	result, err := divideNums(1, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	a, err := strconv.Atoi("123")

	fmt.Printf("%T", str, a, result)
}

func divideNums(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("we can not divide by 0")
	}

	return a / b, nil
}

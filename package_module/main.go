package main

import (
	"errors"
	"fmt"
	"log"
	"modul/constants"
	"strconv"
)

func main() {
	n := "10"

	s, err := strconv.Atoi(n)
	if err != nil {
		log.Println("error while converting string to int", err)
		return
	}

	fmt.Printf("%T", s)

	nums := []int{12, 30, 0, 60}

	for _, num := range nums {
		result, err := div(num)
		if err != nil {
			if errors.Is(err, constants.ErrDivisionByZero) {
				fmt.Println("IT IS TRUE")
			}
			fmt.Println("error while dividing 60 by nums one by one", err)
			return
		}

		fmt.Println(result)
	}

}

func div(n int) (int, error) {
	if n == 0 {
		return 0, constants.ErrDivisionByZero
	}
	return 60 / n, nil

}

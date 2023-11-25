package main

import (
	"math/rand"
	"time"
)

func generateRandomNumber(min, max int) (int, float32) {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min, 0.0
}

func main() {

	/*if x, f := generateRandomNumber(3, 6); x == 4 || f == 0.0 {
		fmt.Println("random number is four")
	} else if x == 5 {
		fmt.Println("random number is five")
	} else {
		fmt.Println(f)
	}

	fmt.Println(f)

	x, f := generateRandomNumber(3, 6)
	if x == 4 || f == 0.0 {
		fmt.Println("random number is four")
	} else if x == 5 {
		fmt.Println("random number is five")
	} else {
		fmt.Println(f)
	}
	fmt.Println(x, f)*/
}

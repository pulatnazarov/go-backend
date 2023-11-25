package main

import "fmt"

func main() {
	var (
		p = 5
		q = 3
	)

	// Relational
	// x == y => false
	// z == y => false
	// z == x => true
	// z >= x => true
	// z > x => false
	// y <= x => false
	// x != z => false
	// x == z => true

	
	// Logic
	// && - false && true
	// || - birinchi true ni qidirish
	// a || b || c || d
	// != - teng bolmasa

	// Bitwise
	z := p & q

	fmt.Println("& z is: ", z)
	
	z = p | q

	fmt.Println("| z is: ", z)

	z = p ^ q

	fmt.Println("^ z is: ", z)

}
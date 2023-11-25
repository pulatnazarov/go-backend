package main

import "fmt"

/*
    *
   * *
  *   *
 *     *
* * * * *
*/

func main() {
	var n = 0
	fmt.Print("n = ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		for j := i; j < n-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			if k == 0 || k == i {
				fmt.Print("*_")
			} else if i == n - 1 {
				fmt.Print("*_")
			} else {
				fmt.Print("  ")
			}
		} 
		fmt.Println()
	}
}

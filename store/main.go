package main

import "fmt"

const (
	StartShopCmd = iota + 1
	FinishShopCmd
)

/* HOMEWORK
Select User type:
1 - Boss
2 - Customer
3 - quit
if 1 {
	Enter password:
	if password {
		Boss login:
		1 - report
		2 - product list
		3 - back
		4 - quit
	} else {
		password is wrong
	}
} else if 2 {
	1 - start shopping
	2 - my basket
	3 - finish
	4 - back
	5 - quit
} else if 3 {
	program should quit
} else {
	not found
}
*/

func main() {
	repo := NewRepository(productList)
	store := NewStore(repo)

	for true {
		var cmd int
		fmt.Print(`
		Enter command:
		1 - Start shopping
		2 - Stop shopping
`)
		fmt.Scan(&cmd)

		switch cmd {
		case StartShopCmd:
			store.printStats()
			store.StartSell()
		case FinishShopCmd:
			return
		default:
			fmt.Println("There is not such kind of command!!!")
		}
	}
}

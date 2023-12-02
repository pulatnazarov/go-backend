package main

import "fmt"

const (
	StartShopCmd = iota + 1
	FinishShopCmd
)

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
			store.printStats()
		case FinishShopCmd:
			return
		}
	}
}

package main

import (
	"fmt"
	"go-backend/store/basket"
	"go-backend/store/customer"
	"go-backend/store/product"
	"go-backend/store/repository"
	"go-backend/store/store"
)

const Password = "123"

// Customer cmd
const (
	StartShopCmd = iota + 1
	FinishShopCmd
)

// Main page cmd
const (
	BossCmd = iota + 1
	CustomerCmd
	UserQuitCmd
)

// Boss cmd
const (
	ReportCmd = iota + 1
	ListCmd
	BackCmd
	QuitCmd
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
	repo := repository.NewRepository(product.List)
	s := store.NewStore(repo)

	for true {
		var userCmd int
		fmt.Print(`
		Enter command:
		1 - Boss
		2 - Customer
		3 - Quit
`)
		fmt.Scan(&userCmd)

		switch userCmd {
		case BossCmd:
			var password string
			fmt.Print("Enter password: ")
			fmt.Scan(&password)
			if password != Password {
				fmt.Println("password is wrong")
				return
			}
			fmt.Println("BOSS LOGIN")

			fmt.Println(`
		1 - Report
		2 - Product List
		3 - Back
		4 - Quit
`)
			var bossCmd int
			fmt.Print("Enter command: ")
			fmt.Scan(&bossCmd)

			switch bossCmd {
			case ReportCmd:
				fmt.Println("report selected by boss")
			case ListCmd:
				fmt.Println("list selected by boss")
			case BackCmd:
				fmt.Println("back selected by boss")
			case QuitCmd:
				fmt.Println("quit selected by boss")
			}
		case CustomerCmd:
			b := basket.NewBasket()
			name, cash := customer.GetCustomerInfo()
			user := customer.NewCustomer(name, cash, b)

			var userCmd int
			fmt.Println(`
			1 - start shopping
			2 - my basket
			3 - finish
			4 - back
			5 - quit
`)
			fmt.Scan(&userCmd)

			switch userCmd {
			case StartShopCmd:
				s.StartSell(user)
			case FinishShopCmd:
				return
			}
		case UserQuitCmd:
			return
		default:
			fmt.Println("There is not such kind of command!!!")
		}
	}
}

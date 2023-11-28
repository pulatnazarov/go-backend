package main

import (
	"fmt"
	"math/rand"
)

const chances = 3

type MySlice []int

func forEach() {

}

func main() {
	var nums []int

	nums = append(nums, 1, 2, 3)
	fmt.Println("nums: ", nums)

	var newNums MySlice

	newNums = append(newNums, 2, 3, 4)
	fmt.Println("newNums: ", newNums)

}

type Game struct {
	RandomNumber int
	Player       Player
}

func (g Game) NewGame(player Player) Game {
	return Game{
		RandomNumber: rand.Intn(10),
		Player:       player,
	}
}

func (g Game) StartGame() {
	fmt.Printf("Welcome %s\n", g.Player.Name)
	fmt.Println("This is guessing game")
	for i := 0; i < g.Player.Chances; i++ {
		var n int
		fmt.Printf("%d chance enter number: ", i+1)
		fmt.Scan(&n)

		if n == g.RandomNumber {
			fmt.Println("You won! you found random number in", i+1, "tries")
			return
		} else {
			fmt.Println("Incorrect")
		}
	}

	fmt.Println("You lost!\nThe random number was", g.RandomNumber)
}

type Player struct {
	Name    string
	Chances int
}

func (p Player) NewPlayer(name string, chances int) Player {
	return Player{
		Name:    name,
		Chances: chances,
	}
}

func getPlayerName() string {
	var (
		name string
	)

	fmt.Print("enter your name: ")
	fmt.Scan(&name)

	return name
}

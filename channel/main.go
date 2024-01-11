package main

import "fmt"

/*
["apple", "banana", "banana", "banana", "apple", "cherry"]
writeWords func receives nothing and returns chan string

removeRepeated func receives chan string, remove new channel apple, banana, apple, cherry
*/

func main() {
	ch1 := writeWords()

	newCh := removeRepeated(ch1)

	for word := range newCh {
		fmt.Print(word, ", ")
	}
	fmt.Println()
}

func writeWords() chan string {
	words := []string{"apple", "banana", "banana", "banana","banana", "banana", "banana","banana", "banana", "banana","banana", "banana", "banana", "apple", "cherry"}
	ch1 := make(chan string)

	go func() {
		defer close(ch1)
		for _, word := range words {
			ch1 <- word
		}
	}()

	return ch1
}

func removeRepeated(ch1 chan string) chan string {
	newCh := make(chan string)

	go func ()  {
		defer close(newCh)
		previous := "" // apple
								//   1       1      0       0       1      1
		for word := range ch1 { // apple, banana, banana, banana, apple, cherry
			if previous != word {
				newCh <- word
			}
			previous = word
		}
	}()

	return newCh
}

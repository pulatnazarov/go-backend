package main

import (
	"fmt"
	"sync"
	"time"
)

var nums = []int{
	1, 2, 3, 4,
}

func main() {
	var mu = sync.Mutex{}

	go func() {
		for i := 5; i < 11; i++ {
			mu.Lock()
			nums = append(nums, i)
			mu.Unlock()
			fmt.Println("first goroutine:", nums, i)
		}
	}()

	go func() {
		for j := 10; j > 4; j-- {
			mu.Lock()
			nums = append(nums, j)
			mu.Unlock()
			fmt.Println("second goroutine:", nums, j)
		}
	}()

	time.Sleep(time.Millisecond * 500)
}

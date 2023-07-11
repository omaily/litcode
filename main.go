package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func generateRandomSlise(size int) []int {
	slise := make([]int, size)
	// rand.Seed(time.Now().UnixNano())
    fmt.Print()
	for i := range slise {
		slise[i] = rand.Intn(9) + 1
	}
	return slise
}

func printSlise(slise []int) {
	for i, v := range slise {
		if i%10 == 0 && i != 0 {
			fmt.Println()
		}
		fmt.Printf("%d : ", v)
	}
	println()
}

func main() {
    // SortAndRemoveDuplicatesInPlace()   
    BestTimetoBuyAndSell()
}

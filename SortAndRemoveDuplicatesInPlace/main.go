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
		slise[i] = rand.Intn(10)
	}
	return slise
}

func removeDuplicates(nums []int) int {
	garbageCursor := len(nums) - 1
	for i, v := range nums {
		if i >= garbageCursor {
			break
		}

        //making sort
        var iterator int  = i 
		for iterator > 0 && nums[iterator] < nums[iterator-1]  { 
			nums[iterator], nums[iterator-1] = nums[iterator-1], nums[iterator]
			iterator--
		}
    
		// making Garbage
		for j := i + 1; j < garbageCursor; j++ {
			// если в garbageCursor элемент равный повторяющемуся
			for nums[garbageCursor] == v {
				garbageCursor--
			}
			if j >= garbageCursor {
				break
			}
			// убираем повторяющиеся элементы
			if nums[j] == v {
				nums[j], nums[garbageCursor] = nums[garbageCursor], nums[j]
				garbageCursor--
			}
		}
	}

	return garbageCursor
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
	var size int = 100
	nums := generateRandomSlise(size)
	removeDuplicates(nums)
	printSlise(nums)
}

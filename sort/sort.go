package sort

import (
	"fmt"
	"math/rand"
	"strconv"
)

func generateRandomSlise(size int) []int {
	slise := make([]int, size)
	fmt.Print()
	for i := range slise {
		slise[i] = rand.Intn(10)
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
	println("---------------------------------------")
}

func Sort() {
	// var size int = 15
	// nums := generateRandomSlise(size)

	n := 6
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, fiftine())
		} else if i%3 == 0 {
			result = append(result, thri())
		} else if i%5 == 0 {
			result = append(result, five())
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	fmt.Printf("%v\n", result)
}

func thri() string {
	return "Fizz"
}

func five() string {
	return "Buzz"
}

func fiftine() string {
	return "FizzBuzz"
}

package array

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

func linearSearchContains(slise []int, elem int) bool {
	for i := range slise {
		if elem == slise[i] {
			return true
		}
	}
	return false
}

// возвращает первый индекс первого встреченного элемента
// если элемент не найден вернет отрицательное значение (заменяет Contains)
func linearSearchFind(slise []int, elem int) int {
	for i, v := range slise {
		if elem == v {
			return i
		}
	}
	return -1
}

func StartTask() {
	// task_26()
	// task_122()
	// task_189()
	// task_217()
	// task_136()
	// task_350()
	task_66()
}

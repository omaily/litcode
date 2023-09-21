package matrix

import (
	"fmt"
	"math/rand"
)

func createMatrix(row int, col int) [][]int {
	mat := make([][]int, row)
	for i := range mat {
		mat[i] = make([]int, col)
		for j := range mat[i] {
			mat[i][j] = rand.Intn(30)
		}
	}

	return mat
}

func printMatrix(m [][]int) {
	for i := range m {
		for j := range m[i] {
			fmt.Printf("%3d ", m[i][j])
		}
		fmt.Println()
	}
	for i := 0; i < len(m[0]); i++ {
		fmt.Print("")
	}
	fmt.Println()
}

func StartTask() {
	m := createMatrix(4, 4)
	printMatrix(m)

	printMatrix(task_48(m))
}

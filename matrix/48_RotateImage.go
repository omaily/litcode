package matrix

import "fmt"

// 48. Rotate Image
func task_48(m [][]int) [][]int {
	n := len(m) - 1

	for i := 0; i < n-i; i++ {
		for j := i; j < n-i; j++ {
			temp := m[i][j]
			fmt.Printf("[%d,%d]=%d", i, j, temp)
			m[i][j] = m[n-j][i]
			fmt.Printf("[%d,%d]=%d", n-j, i, m[n-j][i])
			m[n-j][i] = m[n-i][n-j]
			fmt.Printf("[%d,%d]=%d", n-j, n-i, m[n-i][n-j])
			m[n-i][n-j] = m[j][n-i]
			fmt.Printf("[%d,%d]=%d", j, n-i, m[j][n-i])
			m[j][n-i] = temp
			fmt.Println()
		}
		printMatrix(m)
	}
	return m
}

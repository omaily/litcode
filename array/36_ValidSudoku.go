package array

import "fmt"

func task_36() {
	// nums := generateRandomSlise(17)
	nums := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Printf("nums1=%v\n", nums)
	isValidSudoku(nums)
}

func isValidSudoku(board [][]byte) bool {

	unsorted_map := make(map[string]bool, 250)
	for i, line := range board {
		for j, cell := range line {
			if cell == '.' {
				continue
			}
			key := fmt.Sprint(i/3*3, j/3, cell)
			keyLine := fmt.Sprint("line", i, cell)
			keyColum := fmt.Sprint("colum", j, cell)
			if unsorted_map[key] || unsorted_map[keyLine] || unsorted_map[keyColum] {
				return false
			}
			unsorted_map[key], unsorted_map[keyLine], unsorted_map[keyColum] = true, true, true
		}
	}

	return true
}

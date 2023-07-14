package array

import "fmt"

func task_66() {
	nums := generateRandomSlise(17)
	fmt.Printf("nums1=%v\n", nums)
	fmt.Printf("itog =%v\n", plusOne([]int{9, 9}))
}

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			if i == 0 {
				incDigits := make([]int, n+1)
				copy(incDigits[1:], digits)
				incDigits[0] = 1
				return incDigits
			}
			continue
		}
		digits[i]++
		return digits
	}
	return digits
}

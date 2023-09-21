package array

import "fmt"

func task_66() {
	nums := generateRandomSlise(17)
	fmt.Printf("nums1=%v\n", nums)
	fmt.Printf("itog =%v\n", plusOne2([]int{9, 9}))
}

func plusOne1(digits []int) []int {
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

func plusOne2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	result := make([]int, 1, len(digits)+1)
	result[0] = 1
	return append(result, digits[:]...)
}

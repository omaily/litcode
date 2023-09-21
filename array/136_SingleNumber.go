package array

import "sort"

func task_136() {
	println(singleNumber([]int{4, 1, 2, 1, 2}))
}

func singleNumberSlow(nums []int) int {
	sort.Ints(nums)
	printSlise(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			return nums[i]
		} else {
			i++
		}
	}
	return nums[len(nums)-1]
}

// best praktis
func singleNumber(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		nums[i+1] = nums[i] ^ nums[i+1]
	}
	return nums[len(nums)-1]
}

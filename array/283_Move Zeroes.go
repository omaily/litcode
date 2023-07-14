package array

import "fmt"

func task_283() {
	nums := generateRandomSlise(17)
	fmt.Printf("nums1=%v\n", nums)
	moveZeroes(nums)
}

func moveZeroes(nums []int) {
	isFirstMeetingZero := false
	var beginLoopZero int
	for i := range nums {
		if nums[i] == 0 && !isFirstMeetingZero {
			isFirstMeetingZero = true
			beginLoopZero = i
		}
		if nums[i] != 0 && isFirstMeetingZero {
			nums[i], nums[beginLoopZero] = nums[beginLoopZero], nums[i]
			beginLoopZero++
		}
	}
	fmt.Printf("itog =%v\n", nums)
}

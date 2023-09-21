package array

import "fmt"

func task_283() {
	nums := generateRandomSlise(17)
	fmt.Printf("nums1=%v\n", nums)
	moveZeroes3(nums)
}

func moveZeroes1(nums []int) {
	isFirstMeetingZero := true
	var beginLoopZero int
	for i := range nums {
		if nums[i] == 0 && isFirstMeetingZero {
			isFirstMeetingZero = false
			beginLoopZero = i
		}
		if nums[i] != 0 && !isFirstMeetingZero {
			nums[i], nums[beginLoopZero] = nums[beginLoopZero], nums[i]
			beginLoopZero++
		}
	}
	fmt.Printf("itog =%v\n", nums)
}

func moveZeroes2(nums []int) {
	var windowZero int
	var windowLen int
	for i := range nums {
		if nums[i] == 0 {
			windowZero = i - windowLen
			windowLen++
			continue
		}
		if windowLen > 0 {
			nums[i], nums[windowZero] = nums[windowZero], nums[i]
			windowZero++
		}
	}
}

func moveZeroes3(nums []int) {
	var windowZero int
	for i := range nums {
		if nums[i] == 0 {
			windowZero++
			continue
		}
		if windowZero > 0 {
			nums[i], nums[i-windowZero] = nums[i-windowZero], nums[i]
		}
	}
}

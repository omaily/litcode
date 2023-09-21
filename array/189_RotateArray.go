package array

func task_189() {
	var size int = 20
	nums := generateRandomSlise(size)
	printSlise(nums)
	rotateInPlase(nums, 4)
	println()
	printSlise(nums)
}

// решение задачи с помощью второго массива
func rotate(nums []int, k int) {
	length := len(nums)
	if length <= 0 && k <= 0 {
		return
	}
	numsCopy := make([]int, length)
	copy(numsCopy, nums)

	for i := range nums {
		nums[(i+k)%length] = numsCopy[i]
	}
}

// решение задачи на месте
func rotateInPlase(nums []int, k int) {
	length := len(nums)

	for i := 0; i < k; i++ {
		lastValue := nums[length-1]
		for j := length - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = lastValue
	}
}

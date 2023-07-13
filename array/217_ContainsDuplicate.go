package array

func task_217() {
	var size int = 10
	nums := generateRandomSlise(size)
	printSlise(nums)
	println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {

				return true
			}
		}
	}
	return false
}

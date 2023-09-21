package array

func task_26() {
	var size int = 100
	nums := generateRandomSlise(size)
	removeDuplicates2(nums)
	printSlise(nums)
}

func removeDuplicates1(nums []int) {
	printSlise(nums)
	var rep int
	for i := 0; i < len(nums); i++ {
		var repet bool
		for j := 0; j < rep; j++ {
			if nums[i] == nums[j] {
				repet = true
				break
			}
		}
		if repet {
			continue
		}
		nums[i], nums[rep] = nums[rep], nums[i]
		rep++
	}

	printSlise(nums)
}

func removeDuplicates2(nums []int) {
	printSlise(nums)
	hash := make(map[int]struct{}, len(nums))
	var iterator int
	for i, v := range nums {
		if _, ok := hash[v]; !ok {
			nums[iterator], nums[i] = nums[i], nums[iterator]
			iterator++
		}
		hash[v] = struct{}{}
	}
	printSlise(nums)
}

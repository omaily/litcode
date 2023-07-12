package array

func SortAndRemoveDuplicatesInPlace() {
	var size int = 100
	nums := generateRandomSlise(size)
	removeDuplicates(nums)
	printSlise(nums)
}

func removeDuplicates(nums []int) int {
	garbageCursor := len(nums) - 1
	for i, v := range nums {
		if i >= garbageCursor {
			break
		}

		//making sort
		var iterator int = i
		for iterator > 0 && nums[iterator] < nums[iterator-1] {
			nums[iterator], nums[iterator-1] = nums[iterator-1], nums[iterator]
			iterator--
		}

		// making Garbage
		for j := i + 1; j < garbageCursor; j++ {
			// если в garbageCursor элемент равный повторяющемуся
			for nums[garbageCursor] == v {
				garbageCursor--
			}
			if j >= garbageCursor {
				break
			}
			// убираем повторяющиеся элементы
			if nums[j] == v {
				nums[j], nums[garbageCursor] = nums[garbageCursor], nums[j]
				garbageCursor--
			}
		}
	}

	return garbageCursor
}

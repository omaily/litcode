package array

import "fmt"

func task_350() {
	nums := generateRandomSlise(17)
	nums2 := generateRandomSlise(7)
	fmt.Printf("nums1=%v\n", nums)
	fmt.Printf("nums2=%v\n", nums2)
	fmt.Printf("itog=%v\n", intersect(nums, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	unsigned_map := make(map[int]int)
	for _, key := range nums1 {
		unsigned_map[key]++
	}
	intersection := make([]int, len(nums2)/4)

	for _, key := range nums2 {
		if unsigned_map[key] > 0 {
			intersection = append(intersection, key)
			unsigned_map[key]--
		}
	}
	return intersection
}

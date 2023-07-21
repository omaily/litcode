package stringi

import "fmt"

func task_387() {
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
}

func firstUniqChar(s string) int {
	unsorted_map := make(map[byte]bool, len(s)/2)
	for i := range s {
		if unsorted_map[s[i]] {
			continue
		}
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				unsorted_map[s[j]] = true
			}
		}
		if !unsorted_map[s[i]] {
			return i
		}
	}
	return -1
}

func firstUniqChar_v2(s string) int {
	unsorted_map := make(map[rune]int)
	for _, v := range s {
		unsorted_map[v]++
	}
	for i, v := range s {
		if unsorted_map[v] == 1 {
			return i
		}
	}
	return -1
}

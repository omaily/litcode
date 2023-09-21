package list

// Merge Two Sorted Lists
func task_21(list1 *ListNode, list2 *ListNode) *ListNode {
	// return iterativ21(list1, list2)
	return recursev21(list1, list2)
}

func iterativ21(list1 *ListNode, list2 *ListNode) *ListNode {

	result := list1
	temp := &ListNode{}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1
			list1 = list1.Next
			temp = temp.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
			temp = temp.Next
		}
	}

	if list1 == nil {
		temp.Next = list2
	}
	if list2 == nil {
		temp.Next = list1
	}
	return result
}

func recursev21(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = recursev21(list1.Next, list2)
		return list1
	}

	if list2.Val < list1.Val {
		list2.Next = recursev21(list1, list2.Next)
		return list2
	}

	return nil
}

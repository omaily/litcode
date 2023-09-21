package list

// Palindrome Linked List
func task_234(head *ListNode) bool {
	// return  lame234(head)
	return best234(head)
}

func lame234(head *ListNode) bool {

	for head != nil {
		temp := head

		for temp.Next != nil && temp.Next.Next != nil {
			temp = temp.Next
		}

		if temp.Next != nil && head.Val != temp.Next.Val {
			return false
		}

		temp.Next = nil
		head = head.Next

	}

	return true
}

func best234(head *ListNode) bool {
	slow, floid := head, head

	for floid != nil && floid.Next != nil {
		slow = slow.Next
		floid = floid.Next.Next
	}

	var prev *ListNode
	for slow != nil {
		temp := slow.Next
		slow.Next = prev
		prev = slow
		slow = temp
	}

	for prev != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}
	return true
}

package list

// Reverse Linked List
func task_206(head *ListNode) *ListNode {
	return iterativ206(head)
	// return recursev206(head)
}

func iterativ206(head *ListNode) *ListNode {
	var revers *ListNode = nil
	for head != nil {
		temp := head
		head = head.Next
		temp.Next = revers
		revers = temp
	}
	return revers
}

func recursev206(head *ListNode) *ListNode {
	return nil
}

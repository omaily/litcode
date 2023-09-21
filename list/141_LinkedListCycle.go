package list

func task_141(head *ListNode) bool {
	fast := head
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next
		if head == fast {
			return true
		}
	}
	return false
}

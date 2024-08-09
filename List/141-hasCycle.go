package List

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	for slow != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

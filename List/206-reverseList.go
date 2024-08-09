package List

func reverseList(head *ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: nil}
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = dummyHead.Next
		dummyHead.Next = curr
		curr = temp
	}
	return dummyHead.Next
}

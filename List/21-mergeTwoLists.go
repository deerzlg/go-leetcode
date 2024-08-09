package List

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{Next: nil}
	curr, curr1, curr2 := dummyHead, list1, list2
	for curr1 != nil && curr2 != nil {
		if curr1.Val < curr2.Val {
			curr.Next = &ListNode{Val: curr1.Val, Next: nil}
			curr = curr.Next
			curr1 = curr1.Next
		} else {
			curr.Next = &ListNode{Val: curr2.Val, Next: nil}
			curr = curr.Next
			curr2 = curr2.Next
		}
	}
	for curr1 != nil {
		curr.Next = &ListNode{Val: curr1.Val, Next: nil}
		curr = curr.Next
		curr1 = curr1.Next
	}
	for curr2 != nil {
		curr.Next = &ListNode{Val: curr2.Val, Next: nil}
		curr = curr.Next
		curr2 = curr2.Next
	}
	return dummyHead.Next
}

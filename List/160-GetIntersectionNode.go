package List

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	indexA, indexB := headA, headB
	for indexA != nil {
		lenA++
		indexA = indexA.Next
	}
	for indexB != nil {
		lenB++
		indexB = indexB.Next
	}
	if lenA < lenB {
		lenA, lenB = lenB, lenA
		headA, headB = headB, headA
	}
	for i := 0; i < lenA-lenB; i++ {
		headA = headA.Next
	}
	for headA != nil && headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headA
}

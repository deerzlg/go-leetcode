package List

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1, Next: nil}
	curr, curr1, curr2 := dummy, l1, l2
	carry := 0
	for curr1 != nil && curr2 != nil {
		sum := carry + curr1.Val + curr2.Val
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		curr1 = curr1.Next
		curr2 = curr2.Next
	}
	if curr2 != nil {
		curr1, curr2 = curr2, curr1
	}
	for curr1 != nil {
		sum := carry + curr1.Val
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		curr1 = curr1.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

// 更优雅写法
func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哨兵节点
	cur := dummy
	carry := 0                                 // 进位
	for l1 != nil || l2 != nil || carry != 0 { // 有一个不是空节点，或者还有进位，就继续迭代
		if l1 != nil {
			carry += l1.Val // 节点值和进位加在一起
		}
		if l2 != nil {
			carry += l2.Val // 节点值和进位加在一起
		}
		cur.Next = &ListNode{Val: carry % 10} // 每个节点保存一个数位
		carry /= 10                           // 新的进位
		cur = cur.Next                        // 下一个节点
		if l1 != nil {
			l1 = l1.Next // 下一个节点
		}
		if l2 != nil {
			l2 = l2.Next // 下一个节点
		}
	}
	return dummy.Next // 哨兵节点的下一个节点就是头节点
}

package List

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//由于可能会删除头部，所以使用虚拟头结点
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i < n; i++ { //快指针先移动n步
		fast = fast.Next
	}
	for fast.Next != nil { //定位到删除元素的前一个元素
		slow = slow.Next
		fast = fast.Next
	}
	// 左指针的下一个节点就是倒数第 n 个节点
	temp := slow.Next
	//golang会周期性的进行垃圾回收，当一个对象不再被程序中的任何部分*引用*时,垃圾回收器会自动将其回收。
	slow.Next = slow.Next.Next
	//删除节点时显式地将其 Next 指针置为 nil，可以帮助回收机制更快识别
	temp.Next = nil
	return dummy.Next
}

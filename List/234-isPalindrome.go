package List

func isPalindrome(head *ListNode) bool {
	nums := make([]int, 0)
	curr := head
	for curr != nil {
		nums = append(nums, curr.Val)
		curr = curr.Next
	}
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-i-1] {
			return false
		}
	}
	return true
}

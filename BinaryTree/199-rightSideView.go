package BinaryTree

import "container/list"

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			if i == length-1 {
				ans = append(ans, node.Val)
			}
		}
	}
	return ans
}

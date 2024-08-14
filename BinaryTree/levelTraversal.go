package BinaryTree

import (
	"container/list"
	"slices"
)

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		n := queue.Len()
		temp := []int{}
		for i := 0; i < n; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans = append(ans, temp)
	}
	return ans
}

func levelOrderBottom(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		tempArr := []int{}
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			tempArr = append(tempArr, node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans = append(ans, tempArr)
	}
	slices.Reverse(ans)
	return ans
}

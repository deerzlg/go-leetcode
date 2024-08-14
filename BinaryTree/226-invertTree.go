package BinaryTree

import "container/list"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
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
			node.Left, node.Right = node.Right, node.Left
		}
	}
	return root
}

func ReInvertTree(root *TreeNode) *TreeNode {
	var postOrderInvert func(node *TreeNode)
	postOrderInvert = func(node *TreeNode) {
		if node == nil {
			return
		}
		postOrderInvert(node.Left)
		postOrderInvert(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}
	postOrderInvert(root)
	return root
}

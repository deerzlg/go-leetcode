package BinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 框架都一样，仅访问位置不同
func RePreOrderTraversal(root *TreeNode) []int {
	ans := []int{}
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return ans
}

func ReInorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return ans
}

func RePostOrderTraversal(root *TreeNode) []int {
	ans := []int{}
	var postorder func(node *TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		ans = append(ans, node.Val)
	}
	postorder(root)
	return ans
}

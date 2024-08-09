package DataStructure

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   any
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{
		Val:   value,
		Left:  nil,
		Right: nil,
	}
}

// 根据给定的数组构建二叉树
func BuildTree(values []interface{}) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	return buildTreeHelper(values, 0)
}

// 辅助函数：递归构建二叉树
func buildTreeHelper(values []interface{}, index int) *TreeNode {
	if index >= len(values) || values[index] == nil {
		return nil
	}
	node := NewTreeNode(values[index].(int))
	node.Left = buildTreeHelper(values, 2*index+1)
	node.Right = buildTreeHelper(values, 2*index+2)
	return node
}

/* 层序遍历 */
func levelOrder(root *TreeNode) []any {
	// 初始化队列，加入根节点
	queue := list.New()
	queue.PushBack(root)
	// 初始化一个切片，用于保存遍历序列
	nums := make([]any, 0)
	for queue.Len() > 0 {
		// 队列出队
		node := queue.Remove(queue.Front()).(*TreeNode)
		// 保存节点值
		nums = append(nums, node.Val)
		if node.Left != nil {
			// 左子节点入队
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			// 右子节点入队
			queue.PushBack(node.Right)
		}
	}
	return nums
}

func TestBinaryTree() {
	RootNode := BuildTree([]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(levelOrder(RootNode))
}

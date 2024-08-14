package BinaryTree

import (
	"container/list"
)

func ItPreOrderTraversal(root *TreeNode) []int {
	ans := []int{}
	st := list.New()
	if root == nil {
		return ans
	}
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}

func ItInorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	st := list.New()
	//中序遍历的访问和处理顺序不一致，所以需要使用指针辅助遍历
	curr := root
	for curr != nil || st.Len() > 0 {
		//指针来访问节点，访问到最底层
		if curr != nil {
			st.PushBack(curr) // 将访问的节点放进栈
			curr = curr.Left  //左
		} else {
			// 从栈里弹出的数据，就是要处理的数据（放进ans数组里的数据）
			curr = st.Remove(st.Back()).(*TreeNode)
			ans = append(ans, curr.Val) //根
			curr = curr.Right           //右
		}
	}
	return ans
}

func ItPostOrderTraversal(root *TreeNode) []int {
	ans := []int{}
	st := list.New()
	if root == nil {
		return ans
	}
	st.PushBack(root)
	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		ans = append(ans, node.Val)
		if node.Left != nil {
			st.PushBack(node.Left)
		}
		if node.Right != nil {
			st.PushBack(node.Right)
		}
	}
	Reverse(ans)
	return ans
}
func Reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}

package BackTrace

import (
	"fmt"
	. "leetcode/DataStructure"
)

/* 判断当前状态是否为解 */
func isSolution(state *[]*TreeNode) bool {
	return len(*state) != 0 && (*state)[len(*state)-1].Val == 7
}

/* 记录解 */
func recordSolution(state *[]*TreeNode, res *[][]*TreeNode) {
	*res = append(*res, append([]*TreeNode{}, *state...))
}

/* 判断在当前状态下，该选择是否合法 */
func isValid(state *[]*TreeNode, choice *TreeNode) bool {
	return choice != nil && choice.Val != 3
}

/* 更新状态 */
func makeChoice(state *[]*TreeNode, choice *TreeNode) {
	*state = append(*state, choice)
}

/* 恢复状态 */
func undoChoice(state *[]*TreeNode, choice *TreeNode) {
	*state = (*state)[:len(*state)-1]
}

/* 回溯算法：例题三 */
func BacktrackIII(state *[]*TreeNode, choices *[]*TreeNode, res *[][]*TreeNode) {
	// 检查是否为解
	if isSolution(state) {
		// 记录解
		recordSolution(state, res)
	}
	// 遍历所有选择
	for _, choice := range *choices {
		// 剪枝：检查选择是否合法
		if isValid(state, choice) {
			// 尝试：做出选择，更新状态
			makeChoice(state, choice)
			// 进行下一轮选择
			temp := make([]*TreeNode, 0)
			temp = append(temp, choice.Left, choice.Right)
			BacktrackIII(state, &temp, res)
			// 回退：撤销选择，恢复到之前的状态
			undoChoice(state, choice)
		}
	}
}

func TestBacktrackIII() {
	RootNode := BuildTree([]interface{}{0, 1, 2, 4, 4, 5, 6, 7, 8})
	res := make([][]*TreeNode, 0)
	state := make([]*TreeNode, 0)
	choices := make([]*TreeNode, 0)
	choices = append(choices, RootNode)
	BacktrackIII(&state, &choices, &res)
	for _, nodes := range res {
		if len(nodes) != 0 {
			for _, node := range nodes {
				fmt.Println(node.Val)
			}
		}
	}
}

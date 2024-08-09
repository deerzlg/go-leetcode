package BackTrace

/* 回溯算法框架

func backtrack(state *State, choices []Choice, res *[]State) {
	// 判断是否为解
	if isSolution(state) {
		// 记录解
		recordSolution(state, res)
		// 不再继续搜索
		return
	}
	// 遍历所有选择
	for _, choice := range choices {
		// 剪枝：判断选择是否合法
		if isValid(state, choice) {
			// 尝试：做出选择，更新状态
			makeChoice(state, choice)
			backtrack(state, choices, res)
			// 回退：撤销选择，恢复到之前的状态
			undoChoice(state, choice)
		}
	}
}

*/

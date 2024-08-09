package Hash

import "fmt"

func longestConsecutive(nums []int) int {
	// 使用 map 去重并存储所有的数
	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	longestLength := 0

	// 只遍历可能是最长序列起始点的元素
	for num := range numSet {
		// 如果 num-1 存在，则 num 不是序列的起始点，跳过
		if _, found := numSet[num-1]; found {
			continue
		}

		// 找到以 num 开头的最长连续序列
		currentNum := num
		currentLength := 1

		//用while true循环加break实现找连续序列
		for {
			if _, found := numSet[currentNum+1]; !found {
				break
			}
			currentNum++
			currentLength++
		}

		if currentLength > longestLength {
			longestLength = currentLength
		}
	}

	return longestLength
}

func TestLongestConsecutive() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}

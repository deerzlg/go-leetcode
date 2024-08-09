package Hash

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		if v, ok := m[target-num]; ok {
			return []int{v, index}
		} else {
			m[num] = index
		}
	}
	return []int{}
}

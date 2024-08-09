package Sort

func BubbleSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		flag := false
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				flag = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if flag == false {
			break
		}
	}
}

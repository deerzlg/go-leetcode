package Sort

func midOfThree(nums []int, left, right int) int {
	vLeft, vMid, vRight := nums[left], nums[(left+right)/2], nums[right]
	if (vLeft >= vMid && vLeft <= vRight) || (vLeft <= vMid && vLeft >= vRight) {
		return left
	} else if (vRight >= vMid && vRight <= vLeft) || (vRight <= vMid && vRight >= vLeft) {
		return right
	}
	return (left + right) / 2
}

func Partition(nums []int, left, right int) int {
	mid := midOfThree(nums, left, right)
	nums[left], nums[mid] = nums[mid], nums[left]
	i, j := left, right
	for i < j {
		//注意：要先j--再i++，因为使用left作为pivot，基准值在左，就应该从右侧开始找
		//不然如果先i++，左边的元素先与基准值进行比较和交换，这样可能会导致一些大于基准值的元素被错误地放在基准值的左边，从而无法正确排序
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func QuickSort(nums []int, left, right int) {
	//子数组长度为1终止递归
	if left >= right {
		return
	}
	pivot := Partition(nums, left, right)
	QuickSort(nums, left, pivot-1)
	QuickSort(nums, pivot+1, right)
}

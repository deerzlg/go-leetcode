package DoublePointer

import "fmt"

func trap(height []int) int {
	left, right, preMax, sufMax, ans := 0, len(height)-1, 0, 0, 0
	for left < right {
		preMax = max(preMax, height[left])
		sufMax = max(sufMax, height[right])
		//能接到的雨水量取决于该位置左右两边的最大高度中的较小值减去当前位置的高度
		if preMax < sufMax {
			ans += preMax - height[left]
			left++
		} else {
			//相等的情况也进入，移动谁都一样
			ans += sufMax - height[right]
			right--
		}
	}
	return ans
}

func TestTrap() {
	height1 := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height1))
	height2 := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height2))
}

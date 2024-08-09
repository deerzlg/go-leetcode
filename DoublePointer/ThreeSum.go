package DoublePointer

import (
	"fmt"
	"slices"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	slices.Sort(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return ans
}

func TestThreeSum() {
	nums1 := []int{-1, 0, 1, 2, -1, -4}
	nums2 := []int{0, 1, 1}
	nums3 := []int{0, 0, 0}
	fmt.Println(threeSum(nums1))
	fmt.Println(threeSum(nums2))
	fmt.Println(threeSum(nums3))
}

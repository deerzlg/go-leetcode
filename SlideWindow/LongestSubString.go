package SlideWindow

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left, maxLen := 0, 0
	window := make(map[rune]struct{})
	for _, char := range s {
		_, ok := window[char]
		for ok {
			delete(window, rune(s[left]))
			left++
			_, ok = window[char]
		}
		window[char] = struct{}{}
		maxLen = max(maxLen, len(window))
	}
	return maxLen
}

func TestLongestSubString() {
	str1 := "abcabcbb"
	str2 := "bbbbb"
	str3 := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(str1))
	fmt.Println(lengthOfLongestSubstring(str2))
	fmt.Println(lengthOfLongestSubstring(str3))
}

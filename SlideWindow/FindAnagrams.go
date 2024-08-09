package SlideWindow

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	pCount, windowCount, ans := [26]int{}, [26]int{}, []int{}
	for i, c := range p {
		pCount[c-'a']++
		windowCount[s[i]-'a']++
	}
	if pCount == windowCount {
		ans = append(ans, 0)
	}
	for right := len(p); right < len(s); right++ {
		windowCount[s[right]-'a']++
		windowCount[s[right-len(p)]-'a']--
		if pCount == windowCount {
			ans = append(ans, right-len(p)+1)
		}
	}
	return ans
}

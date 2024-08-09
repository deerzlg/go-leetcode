package Hash

import (
	"fmt"
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	result := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		slices.Sort(b)
		sortedStr := string(b)
		result[sortedStr] = append(result[sortedStr], str)
	}
	resultSlice := [][]string{}
	for _, value := range result {
		resultSlice = append(resultSlice, value)
	}
	return resultSlice
}

func TestGroupAnagrams() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}

package leetcode

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var minIndex int = 0
	found := false
	for i := 0; i < len(strs); i++ {
		if len(strs[minIndex]) < len(strs[i]) {
			minIndex = i
		}
	}

	for i := len(strs[minIndex]); i >= 0; i-- {
		found = true
		str := strs[minIndex][0:i]
		for j := 0; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], str) {
				found = false
			}
		}
		if found {
			return strs[minIndex][:i]
		}
	}
	return strs[minIndex][0:0]

}

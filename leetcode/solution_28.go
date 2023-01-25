package leetcode

import "strings"

func StrStr(haystack string, needle string) int {
	return strStr(haystack, needle)
}
func strStr(haystack string, needle string) int {
	if len(haystack) < len(haystack) {
		return -1
	}
	var ptr2 int = 0
	var i int = 0
	for ; i < len(haystack) && ptr2 < len(needle); i++ {
		if haystack[i] != needle[ptr2] {
			if ptr2 > 0 {
				i = i - ptr2 + 1
				ptr2 = 0
			}
		}
		if haystack[i] == needle[ptr2] {
			ptr2++
		}
	}
	if ptr2 == len(needle) {
		return i - ptr2
	} else {
		return -1
	}
}

func strStr2(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func strStr3(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

package leetcode

func FirstUniqChar(s string) int {
	return firstUniqChar(s)
}
func firstUniqChar(s string) int {
	var hash []byte = make([]byte, 26)
	var ptr byte
	for i := 0; i < len(s); i++ {
		ptr = s[i] - 'a'
		if hash[ptr] < 3 {
			hash[ptr] += 1
		}
	}
	for j := 0; j < len(s); j++ {
		if hash[s[j]-'a'] == 1 {
			return j
		}
	}
	return -1
}

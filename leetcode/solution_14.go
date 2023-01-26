package leetcode

func longestCommonPrefix(strs []string) string {
	var c1, c2 byte
	for i := 0; i < len(strs[0]); i++ {
		c1 = strs[0][i]
		for j := 0; j < len(strs); j++ {
			if len(strs[j])-1 < i {
				return strs[j]
			}
			c2 = strs[j][i]
			if c1 != c2 {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

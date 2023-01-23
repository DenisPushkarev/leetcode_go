package leetcode

func IsAnagram(s string, t string) bool {
	return isAnagram(s, t)
}

func isAnagram(s string, t string) bool {
	var hash []int16 = make([]int16, 26)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		hash[s[i]-'a'] += 1
	}
	for i := 0; i < len(t); i++ {
		if hash[t[i]-'a'] == 0 {
			return false
		} else {
			hash[t[i]-'a'] -= 1
		}
	}
	for i := 0; i < len(hash); i++ {
		if hash[i] > 0 {
			return false
		}
	}
	return true
}

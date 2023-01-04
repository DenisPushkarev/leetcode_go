package leetcode

func heightChecker(heights []int) int {
	freq := make([]int, 101)
	tmp := 0
	cnt := 0
	for i := 0; i < len(heights); i++ {
		freq[heights[i]]++
	}
	for i := 0; i < len(heights); i++ {
		for ; freq[tmp] == 0; tmp++ {
		}
		if heights[i] != tmp {
			cnt++
		}
		freq[tmp]--
	}
	return cnt
}

func heightChecker_slow(heights []int) int {
	dst := make([]int, len(heights))
	copy(dst, heights)
	for i := 0; i < len(dst); i++ {
		for j := i; j < len(dst); j++ {
			if dst[i] > dst[j] {
				dst[i], dst[j] = dst[j], dst[i]
			}
		}
	}
	tmp := 0
	for i := 0; i < len(heights); i++ {
		if dst[i] != heights[i] {
			tmp++
		}
	}
	return tmp
}

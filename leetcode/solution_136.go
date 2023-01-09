package leetcode

func SingleNumber(nums []int) int {
	return singleNumber(nums)
}

func singleNumber(nums []int) int {
	intSet := make(map[int]struct{})
	var exists = struct{}{}
	for _, n := range nums {
		if _, n1 := intSet[n]; n1 {
			delete(intSet, n)
		} else {
			intSet[n] = exists
		}
		// fmt.Printf("%v, intSet %v\n", v, intSet)
	}
	for n := range intSet {
		return n
	}
	return 0
}

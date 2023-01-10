package leetcode

func twoSum(nums []int, target int) []int {
	var hash map[int]int = make(map[int]int)
	var n2 int
	for i, v := range nums {
		n2 = target - v
		if val, ok := hash[n2]; ok {
			return []int{i, val}
		} else {
			hash[v] = i
		}
	}
	return []int{}
}

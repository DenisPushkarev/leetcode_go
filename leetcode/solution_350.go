package leetcode

import "sort"

func Intersect(nums1 []int, nums2 []int) []int {
	return intersect(nums1, nums2)
}

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	idx := 0
	for _, v := range nums1 {
		for j := idx; j < len(nums2) && nums2[j] <= v; j++ {
			idx = j + 1
			if nums2[j] == v {
				result = append(result, v)
				idx = j + 1
				break
			}
		}
	}
	return result
}

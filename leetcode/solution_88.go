package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr1 := m - 1
	ptr2 := n - 1
	for i := n + m - 1; i >= 0; i-- {
		switch {
		case ptr1 >= 0 && ptr2 >= 0:
			if nums1[ptr1] > nums2[ptr2] {
				nums1[i] = nums1[ptr1]
				ptr1--
			} else {
				nums1[i] = nums2[ptr2]
				ptr2--
			}
		case ptr1 >= 0:
			nums1[i] = nums1[ptr1]
			ptr1--
		case ptr2 >= 0:
			nums1[i] = nums2[ptr2]
			ptr2--
		}
	}
}

package leetcode

func FindNumbers(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		for j := 10; nums[i]/j > 0; j = j * 100 {
			if nums[i]/j < 10 {
				cnt++
			}
		}
	}
	return cnt
}

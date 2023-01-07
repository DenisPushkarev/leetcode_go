package leetcode

func Rotate(nums []int, k int) {
	rotate(nums, k)
}
func rotate(nums []int, k int) {
	l := len(nums)
	if l < 2 || k == 0 || l == k {
		return
	}
	if k > l {
		k = k % l
	}
	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
}

func reverse(arr []int, start int, finish int) {
	for i, k := start, 0; i <= start+(finish-start)/2; i++ {
		arr[i], arr[finish-k] = arr[finish-k], arr[i]
		k++
	}
}

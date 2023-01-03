package leetcode

func replaceElements(arr []int) []int {
	l := len(arr)
	max_val := arr[l-1]
	arr[l-1] = -1
	var tmp int
	for i := l - 2; i >= 0; i-- {
		tmp = arr[i]
		arr[i] = max_val
		if tmp > max_val {
			max_val = tmp
		}
	}
	return arr
}

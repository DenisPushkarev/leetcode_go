package leetcode

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	if arr[0] > arr[1] || arr[len(arr)-1] > arr[len(arr)-2] {
		return false
	}
	prev := arr[0]
	up := true
	for i := 1; i < len(arr); i++ {
		switch {
		case up && prev > arr[i]:
			up = false
		case !up && prev < arr[i]:
			return false
		case prev == arr[i]:
			return false
		}
		prev = arr[i]
	}
	return !up
}

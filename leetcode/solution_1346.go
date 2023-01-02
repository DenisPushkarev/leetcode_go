package leetcode

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			n := arr[i] / 2
			for j := 0; j < len(arr); j++ {
				if i != j && n == arr[j] {
					return true
				}
			}
		}
	}
	return false
}

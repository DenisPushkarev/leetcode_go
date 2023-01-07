package array

func Equal(arr1 []int, arr2 []int) bool {
	eq := true
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			eq = false
		}
	}
	return eq
}

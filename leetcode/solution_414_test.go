package leetcode

import "testing"

func TestThirdMax(t *testing.T) {
	arr, r := []int{3, 2, 1}, 1
	result := thirdMax(arr)
	if result != r {
		t.Logf("Expected %v, but got %v", r, result)
		t.Fail()
	}

	arr, r = []int{1, 2}, 2
	result = thirdMax(arr)
	if result != r {
		t.Logf("Expected %v, but got %v", r, result)
		t.Fail()
	}

	arr, r = []int{2, 2, 3, 1}, 1
	result = thirdMax(arr)
	if result != r {
		t.Logf("Expected %v, but got %v", r, result)
		t.Fail()
	}

	arr, r = []int{-2, -2, -3, -1}, -3
	result = thirdMax(arr)
	if result != r {
		t.Logf("Expected %v, but got %v", r, result)
		t.Fail()
	}

	arr, r = []int{1, 2, -2147483648}, -2147483648
	result = thirdMax(arr)
	if result != r {
		t.Logf("Expected %v, but got %v", r, result)
		t.Fail()
	}

	arr, r = []int{1, 2, -2147483648, -1, -3}, -1
	result = thirdMax(arr)
	if result != r {
		t.Logf("Expected %v, but got %v", r, result)
		t.Fail()
	}

	arr, r = []int{1, 1, 2}, 2
	result = thirdMax(arr)
	if result != r {
		t.Logf("Expected %v, but got %v", r, result)
		t.Fail()
	}

}

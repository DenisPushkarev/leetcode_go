package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	result := removeElement([]int{3, 2, 2, 3}, 3)
	if result != 2 {
		t.Logf("Expected 2, got %v", result)
		t.Fail()
	}

	result = removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	if result != 5 {
		t.Logf("Expected 5, got %v", result)
		t.Fail()
	}
	result = removeElement([]int{0, 1, 2, 2, 3, 2, 4, 2}, 2)
	// nums = 0, 1, 4, 3, | 3, 2, 4, 2
	if result != 4 {
		t.Logf("Expected 4, got %v", result)
		t.Fail()
	}
}
